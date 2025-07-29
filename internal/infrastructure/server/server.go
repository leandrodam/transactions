package server

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/leandrodam/transactions/internal/infrastructure/config"
	"github.com/leandrodam/transactions/internal/infrastructure/validator"
)

type Server struct {
	config     *config.Config
	httpServer *echo.Echo
}

func NewServer(echo *echo.Echo) *Server {
	return &Server{
		config:     config.Load(),
		httpServer: echo,
	}
}

func (s *Server) Start() {
	s.httpServer.Use(middleware.Recover())
	s.httpServer.Use(middleware.CORS())
	s.httpServer.Pre(middleware.RemoveTrailingSlash())
	s.httpServer.Validator = validator.NewValidator()

	resources := NewResources(s.config.Mysql)
	services := NewServices(resources)
	useCases := NewUseCases(services)
	handlers := NewHandlers(useCases)

	s.registerRoutes(handlers)

	s.run()
}

func (s *Server) run() {
	log.Printf("starting server on port %s", s.config.App.Port)

	if err := s.httpServer.Start(":" + s.config.App.Port); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}
