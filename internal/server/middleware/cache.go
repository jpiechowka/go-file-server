package middleware

import "github.com/gofiber/fiber/v2"

func DisableCache() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		ctx.Set(fiber.HeaderExpires, "0")
		ctx.Set(fiber.HeaderPragma, "no-cache")
		ctx.Set(fiber.HeaderCacheControl, "no-store")

		return ctx.Next()
	}
}
