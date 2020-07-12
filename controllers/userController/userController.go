package userController

import (
	"github.com/gofiber/fiber"
)

func GetUser(ctx *fiber.Ctx) {
	ctx.Send("hoooooo")
}
