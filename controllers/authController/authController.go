package authController

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/queries"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/abort"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/jwt"
	"github.com/gofiber/fiber"
)

type user struct {
	ID                   int
	Name                 string
	Email                string
	Photo                string
	Password             string
	PasswordChangedAt    int32
	PasswordResetToken   string
	PasswordResetExpires int32
}

func createSendToken(u user, ctx *fiber.Ctx) {
	token, err := jwt.SignToken(u.ID)
	if err != nil {
		abort.Msg(500, "error making token", ctx)
		return
	}

	ctx.Status(201).JSON(&fiber.Map{
		"status": 201,
		"token":  token,
		"id":     u.ID,
		"name":   u.Name,
		"email":  u.Email,
	})
}

func Signup(ctx *fiber.Ctx) {
	// getting the body the right way
	type signupInput struct {
		Name     string `json:"name"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(signupInput)
	ctx.BodyParser(input)
	// saving the user into the db
	results := []user{}
	err := database.DB.Select(&results, queries.InsertUser(input.Name, input.Email, input.Password))
	if err != nil {
		abort.Err(500, err, ctx)
		return
	}
	// create and send the response token
	createSendToken(results[0], ctx)
}
