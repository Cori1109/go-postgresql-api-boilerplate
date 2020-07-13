package routes

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/controllers/userController"

	"github.com/gofiber/fiber"
)

// SetupUserRoute exported
func SetupUserRoute(app *fiber.App) {
	user := app.Group("/api/v1/user")

	// routes
	user.Get("/getMe", userController.GetUser)
	user.Patch("/updateMe", userController.GetUser)
	user.Patch("/updatePassword", userController.GetUser)
	user.Delete("/deleteMe", userController.GetUser)

}
