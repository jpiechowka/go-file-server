package server

import (
	"net/http"
	"time"

	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/jpiechowka/go-file-server/internal/server/middleware"
)

func (s *FiberFileServer) setupRoutingAndMiddleware() {
	// Middleware configuration
	s.fiber.Use(recover.New()) // Default fiber panic recovery middleware

	if s.config.EnableBasicAuth {
		s.fiber.Use(basicauth.New(basicauth.Config{
			Users: map[string]string{
				s.config.BasicAuthUser: s.config.BasicAuthPassword,
			},
		}))
	}

	s.fiber.Use(logger.New(logger.Config{
		Format:     "[${time}] ${status} - ${ip} ${latency} ${method} ${path}\n",
		TimeFormat: "Mon Jan-_2 15:04:05 MST", // https://gobyexample.com/time-formatting-parsing
		TimeZone:   "Local",
	}))

	s.fiber.Use(middleware.DisableCache())
	s.fiber.Use(middleware.AddSecurityHeaders(s.config.EnableTls))

	s.fiber.Use(compress.New(compress.Config{
		Level: compress.Level(s.config.CompressionLevel),
	}))

	s.fiber.Use(limiter.New(limiter.Config{
		Max:        int(s.config.RateLimitPerMinute),
		Expiration: 1 * time.Minute,
	}))

	s.fiber.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir(s.config.ServeDirectoryPath),
		Browse: !s.config.DisableDirectoryListing,
	}))
}
