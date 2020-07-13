package routes

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/controllers/authController"

	"github.com/gofiber/fiber"
)

// SetupUserRoute exported
func SetupAuthRoute(app *fiber.App) {
	auth := app.Group("/api/v1/auth")

	// routes
	auth.Post("/signup", authController.Signup)
	auth.Post("/login", authController.Signup)
	auth.Post("/forgotPassword", authController.Signup)
	auth.Patch("/resetPassword/:token", authController.Signup)

}
