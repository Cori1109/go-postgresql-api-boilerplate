package main

import (
	"fmt"
	"log"
	"os"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/routes"
	"github.com/gofiber/cors"
	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/middleware"
	"github.com/gofiber/helmet"
	"github.com/joho/godotenv"
)

func main() {
	// load dotenv
	err := godotenv.Load(".env")
	if err != nil {
		fmt.Print("Error loading .env file")
	}
	// Connect to database
	if err := database.Connect(); err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app.Use(middleware.Logger())

	app.Use(cors.New())

	app.Use(helmet.New())

	routes.SetupUserRoute(app)
	routes.SetupAuthRoute(app)

	app.Get("/", func(c *fiber.Ctx) {
		c.Send("Hello, World 👋!")
	})

	app.Listen(os.Getenv("PORT"))
}
