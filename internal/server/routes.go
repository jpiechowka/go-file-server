package server

import (
	"github.com/gofiber/fiber/v2/middleware/basicauth"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/jpiechowka/go-file-server/internal/server/middleware"
	"net/http"
	"time"
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
	s.fiber.Use(middleware.AddSecurityHeaders(false)) // TODO: HSTS config

	s.fiber.Use(compress.New(compress.Config{
		Level: compress.LevelBestCompression, // TODO: compression level or disable provided in config
	}))

	s.fiber.Use(limiter.New(limiter.Config{
		Max:      int(s.config.RateLimitPerMinute),
		Duration: 1 * time.Minute,
	}))

	s.fiber.Use(filesystem.New(filesystem.Config{
		Root:   http.Dir(s.config.ServeDirectoryPath),
		Browse: true,
	}))
}
