package routes

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/controllers/authController"

	"github.com/gofiber/fiber"
)

// SetupUserRoute exported
func SetupAuthRoute(app *fiber.App) {
	auth := app.Group("/api/auth")

	// routes
	auth.Post("/signup", authController.Signup)
	auth.Post("/login", authController.Login)
	auth.Post("/forgotPassword", authController.ForgotPassword)
	auth.Patch("/resetPassword/:token", authController.ResetPassword)

}
