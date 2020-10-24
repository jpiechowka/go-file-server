package server

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jpiechowka/go-file-server/internal/config"
	"time"
)

type FiberFileServer struct {
	config *config.ServerConfig
	fiber  *fiber.App
}

func NewFiberFileServer(config *config.ServerConfig) *FiberFileServer {
	return &FiberFileServer{
		config: config,
		fiber: fiber.New(fiber.Config{
			IdleTimeout: 60 * time.Second,
			GETOnly:     true,
		}),
	}
}

func (s *FiberFileServer) ConfigureAndStart() error {
	s.setupRoutingAndMiddleware()
	return s.fiber.Listen(s.config.Address)
}
