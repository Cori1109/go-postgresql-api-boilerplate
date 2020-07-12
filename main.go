package main

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/routes"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

func main() {
	app := fiber.New()

	app.Use(middleware.Logger())

	app.Use(cors.New())

	app.Use(helmet.New())

	routes.SetupUserRoute(app)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(3001)
}
