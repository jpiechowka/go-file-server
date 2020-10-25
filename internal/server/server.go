package server

import (
	"crypto/tls"
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
			IdleTimeout:           60 * time.Second,
			DisableStartupMessage: true,
			GETOnly:               true,
		}),
	}
}

func (s *FiberFileServer) ConfigureAndStart() error {
	s.setupRoutingAndMiddleware()

	if s.config.EnableTls {
		cer, err := tls.LoadX509KeyPair(s.config.CertFilePath, s.config.KeyFilePath)
		if err != nil {
			return err
		}

		tlsConfig := &tls.Config{Certificates: []tls.Certificate{cer}}

		listener, err := tls.Listen("tcp", s.config.Address, tlsConfig)
		if err != nil {
			return err
		}

		return s.fiber.Listener(listener)
	}

	return s.fiber.Listen(s.config.Address)
}
