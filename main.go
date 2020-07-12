package main

import (
	"log"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/config"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/routes"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
)

func main() {
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	app.Use(cors.New())

	app.Use(helmet.New())

	routes.SetupUserRoute(app)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(config.Dot("PORT"))
}
