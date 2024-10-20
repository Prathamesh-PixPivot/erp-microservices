package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Nerzal/gocloak/v12"
	"github.com/spf13/viper"
)

type KeycloakRepository struct {
	client   gocloak.GoCloak
	token    *gocloak.JWT
	realm    string
	clientID string
}

func NewKeycloakRepository() (*KeycloakRepository, error) {
	client := gocloak.NewClient(viper.GetString("KEYCLOAK_URL"))
	ctx := context.Background()

	// Authenticate as admin to get the token
	token, err := client.LoginAdmin(ctx, viper.GetString("KEYCLOAK_ADMIN_USER"), viper.GetString("KEYCLOAK_ADMIN_PASSWORD"), viper.GetString("KEYCLOAK_ADMIN_REALM"))
	if err != nil {
		return nil, fmt.Errorf("failed to login as admin: %v", err)
	}

	return &KeycloakRepository{
		client:   *client,
		token:    token,
		realm:    viper.GetString("KEYCLOAK_REALM"),
		clientID: viper.GetString("KEYCLOAK_CLIENT_ID"),
	}, nil
}

// CreateUserInKeycloak creates a user in Keycloak using the gocloak library
func (r *KeycloakRepository) CreateUserInKeycloak(email, password, firstName, lastName string) error {
	token, err := r.client.LoginAdmin(context.Background(),
		viper.GetString("KEYCLOAK_ADMIN_USER"),
		viper.GetString("KEYCLOAK_ADMIN_PASSWORD"),
		viper.GetString("KEYCLOAK_REALM"),
	)
	if err != nil {
		log.Printf("Failed to login as admin to Keycloak: %v", err)
		return err
	}

	user := gocloak.User{
		Username:      gocloak.StringP(email),
		Enabled:       gocloak.BoolP(true),
		EmailVerified: gocloak.BoolP(false), // Email not verified by default
		FirstName:     gocloak.StringP(firstName),
		LastName:      gocloak.StringP(lastName),
		Email:         gocloak.StringP(email),
		Credentials: &[]gocloak.CredentialRepresentation{
			{
				Type:      gocloak.StringP("password"),
				Value:     gocloak.StringP(password),
				Temporary: gocloak.BoolP(false),
			},
		},
	}

	userID, err := r.client.CreateUser(context.Background(), token.AccessToken, viper.GetString("KEYCLOAK_REALM"), user)
	if err != nil {
		log.Printf("Failed to create user in Keycloak: %v", err)
		return err
	}

	// Send verification email
	emailParams := gocloak.ExecuteActionsEmail{
		UserID:   gocloak.StringP(userID),
		Actions:  &[]string{"VERIFY_EMAIL"},
		ClientID: gocloak.StringP(viper.GetString("KEYCLOAK_CLIENT_ID")),
	}
	err = r.client.ExecuteActionsEmail(context.Background(), token.AccessToken, viper.GetString("KEYCLOAK_REALM"), emailParams)
	if err != nil {
		log.Printf("Failed to send verification email: %v", err)
		return err
	}

	return nil
}

// AuthenticateUser authenticates a user with Keycloak and returns access/refresh tokens
func (r *KeycloakRepository) AuthenticateUser(email, password string) (*gocloak.JWT, error) {
	token, err := r.client.Login(context.Background(), r.clientID, viper.GetString("KEYCLOAK_CLIENT_SECRET"), r.realm, email, password, "openid")
	if err != nil {
		return nil, errors.New("failed to authenticate user with Keycloak")
	}
	return token, nil
}
