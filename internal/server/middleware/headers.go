package middleware

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

const (
	cspPolicy  = "default-src 'self'"
	hstsMaxAge = 60 * 60 * 24 * 365 * 3 // 3 years
)

func AddSecurityHeaders(enableHsts bool) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderXFrameOptions, "DENY")
		ctx.Set(fiber.HeaderXContentTypeOptions, "nosniff")
		ctx.Set(fiber.HeaderReferrerPolicy, "no-referrer")

		// CSP
		ctx.Set(fiber.HeaderContentSecurityPolicy, cspPolicy)

		// HSTS
		if enableHsts {
			hstsValue := fmt.Sprintf("max-age=%d ;includeSubDomains; preload", hstsMaxAge)
			ctx.Set(fiber.HeaderStrictTransportSecurity, hstsValue)
		}

		return ctx.Next()
	}
}
