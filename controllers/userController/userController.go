package userController

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/models"
	"github.com/gofiber/fiber"
)

func GetMe(ctx *fiber.Ctx) {
	u := ctx.Locals("user").(models.User)
	ctx.Status(200).JSON(&fiber.Map{
		"status": 200,
		"id":     u.ID,
		"name":   u.Name,
		"email":  u.Email,
	})
}
