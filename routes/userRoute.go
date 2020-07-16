package routes

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/controllers/authController"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/controllers/userController"

	"github.com/gofiber/fiber"
)

// SetupUserRoute exported
func SetupUserRoute(app *fiber.App) {
	user := app.Group("/api/user", authController.Protect)

	// routes
	user.Get("/getMe", userController.GetMe)
	user.Patch("/updateMe", userController.UpdateMe)
	user.Patch("/updatePassword", authController.UpdatePassword)
	user.Delete("/deleteMe", userController.DeleteMe)

}
