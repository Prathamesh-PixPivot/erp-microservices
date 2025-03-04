package handler

import (
	"auth-service/grpc/authpb"
	"auth-service/grpc/organizationpb"
	"auth-service/grpc/userpb"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Nerzal/gocloak/v12"
	"github.com/gofiber/fiber/v2"
	"github.com/segmentio/kafka-go"
	"github.com/spf13/viper"
	"golang.org/x/crypto/bcrypt"
)

type AuthHandler struct {
	userClient         userpb.UserServiceClient
	organizationClient organizationpb.OrganizationServiceClient
	keycloakClient     gocloak.GoCloak
	keycloakToken      string
	realm              string
}

func NewAuthHandler(userClient userpb.UserServiceClient, organizationClient organizationpb.OrganizationServiceClient, keycloakClient *gocloak.GoCloak, keycloakToken string, realm string) *AuthHandler {
	return &AuthHandler{
		userClient:         userClient,
		organizationClient: organizationClient,
		keycloakClient:     *keycloakClient,
		keycloakToken:      keycloakToken,
		realm:              realm,
	}
}

// ProduceEvent is a helper function to send events to Kafka
func ProduceEvent(topic, message string) {
	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{"kafka:9092"}, // Adjust the Kafka broker address
		Topic:   topic,
	})

	defer writer.Close()

	err := writer.WriteMessages(context.Background(),
		kafka.Message{Value: []byte(message)},
	)
	if err != nil {
		log.Fatalf("Failed to produce event: %v", err)
	}
}

func createUserInKeycloak(accessToken, realm string, user map[string]interface{}) error {
	url := fmt.Sprintf("http://host.docker.internal:8080/admin/realms/%s/users", realm)
	userJSON, err := json.Marshal(user)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(userJSON))
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("request to Keycloak failed: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("failed to create user in Keycloak, status code: %d", resp.StatusCode)
	}

	log.Println("User successfully created in Keycloak")
	return nil
}

// SignupUser handles the user signup process
func (h *AuthHandler) SignupUser(c *fiber.Ctx) error {
	var req authpb.SignupRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Step 1: Create user in Keycloak
	user := gocloak.User{
		Username: gocloak.StringP(req.Email),
		Email:    gocloak.StringP(req.Email),
		Enabled:  gocloak.BoolP(true),
		Credentials: &[]gocloak.CredentialRepresentation{{
			Type:      gocloak.StringP("password"),
			Value:     gocloak.StringP(req.Password),
			Temporary: gocloak.BoolP(false),
		}},
	}

	userMap := map[string]interface{}{
		"username": user.Username,
		"email":    user.Email,
		"enabled":  user.Enabled,
		"credentials": []map[string]interface{}{
			{
				"type":      *(*user.Credentials)[0].Type,
				"value":     *(*user.Credentials)[0].Value,
				"temporary": *(*user.Credentials)[0].Temporary,
			},
		},
	}
	err := createUserInKeycloak(h.keycloakToken, h.realm, userMap)

	if err != nil {
		log.Printf("Failed to create user in Keycloak: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user in Keycloak"})
	}

	// Step 2: Hash password before sending to user-service
	hashedPassword, err := hashPassword(req.Password)
	if err != nil {
		log.Printf("Failed to hash password: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user"})
	}

	// Step 3: Create user in user-service via gRPC
	userReq := &userpb.CreateUserRequest{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Email:     req.Email,
		Phone:     req.Phone,
		Password:  hashedPassword,
		Role:      req.Role,
	}
	userRes, err := h.userClient.CreateUser(context.Background(), userReq)
	if err != nil {
		log.Printf("Failed to create user in user-service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create user in user-service"})
	}

	// Step 4: Create organization in organization-service via gRPC
	orgReq := &organizationpb.CreateOrganizationRequest{
		Name:    req.OrganizationName,
		GstIn:   req.GstIn,
		Phone:   req.Phone,
		Email:   req.Email,
		Address: req.Address,
		City:    req.City,
		State:   req.State,
		Country: req.Country,
		Zipcode: req.Zipcode,
		Website: req.Website,
	}
	orgRes, err := h.organizationClient.CreateOrganization(context.Background(), orgReq)
	if err != nil {
		log.Printf("Failed to create organization in organization-service: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to create organization in organization-service"})
	}

	// Produce event after successful signup
	eventMessage := fmt.Sprintf(`{"userId": "%d", "email": "%s"}`, userRes.Id, userReq.Email)
	ProduceEvent("user-signup", eventMessage)

	// Return success response
	return c.JSON(fiber.Map{
		"message":         "Signup successful",
		"user_id":         userRes.Id,
		"organization_id": orgRes.Id,
	})
}

// SigninUser handles user login and token generation via Keycloak
func (h *AuthHandler) SigninUser(c *fiber.Ctx) error {
	var req authpb.SigninRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid request"})
	}

	// Authenticate user via Keycloak
	token, err := h.keycloakClient.Login(c.Context(), viper.GetString("KEYCLOAK_CLIENT_ID"), viper.GetString("KEYCLOAK_CLIENT_SECRET"), h.realm, req.Email, req.Password)
	if err != nil {
		log.Printf("Realm: %s\nemail: %s\npassword: %s", h.realm, req.Email, req.Password)
		log.Printf("Failed to authenticate user: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	// Get user information from Keycloak
	userInfo, err := h.keycloakClient.GetUserInfo(c.Context(), token.AccessToken, h.realm)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	// Convert ExpiresIn (int, seconds) to time.Time by adding to current time
	accessTokenExpiration := time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	refreshTokenExpiration := time.Now().Add(time.Duration(token.RefreshExpiresIn) * time.Second)

	c.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    token.AccessToken,
		Expires:  accessTokenExpiration, // Set the converted expiration time for access token
		HTTPOnly: true,                  // Secure the cookie (HTTP-only, no JS access)
		Secure:   true,                  // Use in HTTPS environments
	})

	// Set the new refresh token in the response cookie
	c.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    token.RefreshToken,
		Expires:  refreshTokenExpiration, // Set the converted expiration time for refresh token
		HTTPOnly: true,                   // Secure the cookie (HTTP-only, no JS access)
		Secure:   true,                   // Use in HTTPS environments
	})

	// fmt.Println(token.AccessToken)
	// Return tokens and user details
	// return c.JSON(&authpb.SigninResponse{
	// 	UserId:       *userInfo.Sub,
	// 	FirstName:    *userInfo.GivenName,
	// 	LastName:     *userInfo.FamilyName,
	// 	Email:        *userInfo.Email,
	// 	AccessToken:  token.AccessToken,
	// 	RefreshToken: token.RefreshToken,
	// })
	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":        *userInfo.Sub,
			"firstName": *userInfo.GivenName,
			"lastName":  *userInfo.FamilyName,
			"email":     *userInfo.Email,
		},
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

// RefreshToken handles token refresh via Keycloak
func (h *AuthHandler) RefreshToken(c *fiber.Ctx) error {
	refreshToken := c.Cookies("rt")
	if refreshToken == "" {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"error": "Refresh token is missing or invalid",
		})
	}

	// Refresh token via Keycloak
	token, err := h.keycloakClient.RefreshToken(c.Context(), refreshToken, viper.GetString("KEYCLOAK_CLIENT_ID"), viper.GetString("KEYCLOAK_CLIENT_SECRET"), h.realm)
	if err != nil {
		log.Printf("Failed to refresh token: %v", err)
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid refresh token"})
	}

	// Get user information from Keycloak
	userInfo, err := h.keycloakClient.GetUserInfo(c.Context(), token.AccessToken, h.realm)
	if err != nil {
		log.Printf("Failed to get user info: %v", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "Failed to get user info"})
	}

	// Convert ExpiresIn (int, seconds) to time.Time by adding to current time
	accessTokenExpiration := time.Now().Add(time.Duration(token.ExpiresIn) * time.Second)
	refreshTokenExpiration := time.Now().Add(time.Duration(token.RefreshExpiresIn) * time.Second)

	c.Cookie(&fiber.Cookie{
		Name:     "at",
		Value:    token.AccessToken,
		Expires:  accessTokenExpiration, // Set the converted expiration time for access token
		HTTPOnly: true,                  // Secure the cookie (HTTP-only, no JS access)
		Secure:   true,                  // Use in HTTPS environments
	})

	// Set the new refresh token in the response cookie
	c.Cookie(&fiber.Cookie{
		Name:     "rt",
		Value:    token.RefreshToken,
		Expires:  refreshTokenExpiration, // Set the converted expiration time for refresh token
		HTTPOnly: true,                   // Secure the cookie (HTTP-only, no JS access)
		Secure:   true,                   // Use in HTTPS environments
	})

	// Return new tokens (access token and refresh token)
	return c.JSON(fiber.Map{
		"user": fiber.Map{
			"id":        *userInfo.Sub,
			"firstName": *userInfo.GivenName,
			"lastName":  *userInfo.FamilyName,
			"email":     *userInfo.Email,
		},
		"access_token":  token.AccessToken,
		"refresh_token": token.RefreshToken,
	})
}

// hashPassword hashes the given password using bcrypt.
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}
