package abort

import "github.com/gofiber/fiber"

func Msg(status int, msg string, ctx *fiber.Ctx) {
	ctx.Status(status).JSON(&fiber.Map{
		"status":  status,
		"message": msg,
	})
}

func Err(status int, err error, ctx *fiber.Ctx) {
	ctx.Status(status).JSON(&fiber.Map{
		"status": status,
		"err":    err,
	})
}
