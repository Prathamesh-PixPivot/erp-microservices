package main

import (
	"log"
	"net/http"

	service "itom/internal"

	"github.com/gofiber/fiber/v2"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

// Server represents the ITOM server using Fiber.
type Server struct {
	ITOMService *service.ITOMService
	Addr        string
	Config      *service.Config
	App         *fiber.App
}

// convertToHTTPRequest converts a fasthttp.RequestCtx to an http.Request.
func convertToHTTPRequest(c *fiber.Ctx) (*http.Request, error) {
	req := new(http.Request)
	err := fasthttpadaptor.ConvertRequest(c.Context(), req, true)
	if err != nil {
		return nil, err
	}
	return req, nil
}

// NewFiberServer creates a new Fiber-based ITOM server.
func NewFiberServer(cfg *service.Config, srv *service.ITOMService) *Server {
	app := fiber.New()
	s := &Server{
		ITOMService: srv,
		Addr:        ":" + cfg.Port,
		Config:      cfg,
		App:         app,
	}
	// Register a default route.
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Welcome to IT Operations Management (ITOM) Service")
	})
	// Register the /alert endpoint.
	app.Post("/alert", s.handleAlert)

	return s
}

func (s *Server) handleAlert(c *fiber.Ctx) error {
	// Convert the Fiber context to an HTTP request.
	httpReq, err := convertToHTTPRequest(c)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to convert request",
		})
	}

	// Delegate processing to ITOMService.
	if err := s.ITOMService.ProcessAlert(httpReq); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "failed to process alert",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "alert processed successfully",
	})
}

// Start launches the Fiber server.
func (s *Server) Start() error {
	log.Printf("Starting ITOM server on %s", s.Addr)
	if s.Config.UseTLS {
		return s.App.ListenTLS(s.Addr, s.Config.TLSCertFile, s.Config.TLSKeyFile)
	}
	return s.App.Listen(s.Addr)
}
