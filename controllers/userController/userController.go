package userController

import (
	"fmt"

	"github.com/gofiber/fiber"
)

func GetMe(ctx *fiber.Ctx) {
	fmt.Println("🚨🚨🚨 hey from getMe")
	fmt.Println(ctx.Locals("user"))

	// ctx.Send("getMe")
}
