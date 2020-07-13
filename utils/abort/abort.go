package abort

import "github.com/gofiber/fiber"

func Now(status int, msg string, ctx *fiber.Ctx) {
	ctx.Status(status).JSON(&fiber.Map{
		"status":  status,
		"message": msg,
	})
}
