package authController

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/abort"
	"github.com/gofiber/fiber"
)

// func createSendToken(){

// }

func Signup(ctx *fiber.Ctx) {

	// getting the body the right way
	type signupInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(signupInput)
	ctx.BodyParser(input)

	q := InsertUser(input.Name, input.Email, input.Password)

	// saving the user into the db
	rows, err := database.DB.Query("")
	if err != nil {
		ctx.Status(500).JSON(&fiber.Map{
			"success": false,
			"error":   err,
		})
		return
	}
	defer rows.Close()

	abort.Now(200, "hoooo", ctx)
}
