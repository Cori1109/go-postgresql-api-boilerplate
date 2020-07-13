package authController

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/abort"
	"github.com/gofiber/fiber"
)

func Signup(ctx *fiber.Ctx) {
	abort.Now(200, "hoooo", ctx)
}
