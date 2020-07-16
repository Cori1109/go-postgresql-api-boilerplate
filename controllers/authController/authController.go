package authController

import (
	"fmt"

	"time"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/queries"

	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/abort"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/email"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/jwt"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/password"

	"github.com/gofiber/fiber"
)

type user struct {
	ID                     int
	Name                   string
	Email                  string
	Photo                  string
	Password               string
	Password_changed_at    float64
	Password_reset_token   string
	Password_reset_expires float64
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
	// check if all needed input is here
	if input.Name == "" || input.Email == "" || input.Password == "" {
		abort.Msg(400, "you need to provide email, name and password input.", ctx)
		return
	}
	// hash password
	hashed, hachingErr := password.HashPassword(input.Password)
	if hachingErr != nil {
		abort.Err(500, hachingErr, ctx)
		return
	}
	// saving the user into the db
	results := []user{}
	err := database.DB.Select(&results, queries.InsertUser(input.Name, input.Email, hashed))
	if err != nil {
		abort.Err(500, err, ctx)
		return
	}
	// create and send the response token
	createSendToken(results[0], ctx)
}

func Login(ctx *fiber.Ctx) {
	// getting the body the right way
	type loginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	input := new(loginInput)
	ctx.BodyParser(input)
	// 1) check if all needed input is here
	if input.Email == "" || input.Password == "" {
		abort.Msg(400, "you need to provide email and password input.", ctx)
		return
	}
	//  2) Check if user exists
	result := user{}
	err := database.DB.Get(&result, queries.GetUserWithEmail(input.Email))
	if err != nil {
		abort.Msg(400, "no user with this email", ctx)
		return
	}
	// 3) check if password is correct
	if !password.CheckPasswordHash(result.Password, input.Password) {
		abort.Msg(401, "the password you enterd is wrong, please try again", ctx)
		return
	}
	// 4) If everything is ok
	// create and send the response token
	createSendToken(result, ctx)
}

func ForgotPassword(ctx *fiber.Ctx) {
	// getting the body the right way
	type fpInput struct {
		Email string `json:"email"`
	}
	input := new(fpInput)
	ctx.BodyParser(input)
	// 1) check if all needed input is here
	if input.Email == "" {
		abort.Msg(400, "you need to provide email input.", ctx)
		return
	}
	//  2) Check if user exists
	u := user{}
	err := database.DB.Get(&u, queries.GetUserWithEmail(input.Email))
	if err != nil {
		abort.Msg(400, "no user with this email", ctx)
		return
	}

	resetData := password.CreatePasswordResetToken()

	// 3) save token reset data in db
	_, err2 := database.DB.Exec(queries.UpdateUserPassResetData(u.ID, resetData.Prt, resetData.Pre))
	if err2 != nil {
		abort.Err(400, err2, ctx)
		return
	}
	// 4) Send it to user's email
	url := fmt.Sprintf("http://localhost:3001/api/auth/resetPassword/%s", resetData.Rt)
	html := fmt.Sprintf("<b>LINK RESET YOUR PASSWORD: %s</b>", url)
	fmt.Println("📌", url)
	_, err3 := email.SendOne("reset your password", u.Name, u.Email, url, html)
	if err3 != nil {
		abort.Err(400, err3, ctx)
	}

	abort.Msg(200, "email sent to your inbox", ctx)
}

func ResetPassword(ctx *fiber.Ctx) {
	// 1) Get user based on the token
	token := ctx.Params("token")

	crypted := password.CryptString(token)
	now := time.Now().Unix()

	//  2) Check if user exists
	u := user{}
	err := database.DB.Get(&u, queries.GetUserByResetToken(crypted, now))
	if err != nil {
		abort.Msg(400, "Token is invalid or has expired", ctx)
		return
	}
	// 3) hash an Update password, changedPasswordAt property for the user

	// getting the body the right way
	type rpInput struct {
		NewPassword string `json:"newPassword"`
	}
	input := new(rpInput)
	ctx.BodyParser(input)
	// check if all needed input is here
	if input.NewPassword == "" {
		abort.Msg(400, "you need to provide new password input.", ctx)
		return
	}

	// hashing the newPassword
	hashed, hachingErr := password.HashPassword(input.NewPassword)
	if hachingErr != nil {
		abort.Err(500, hachingErr, ctx)
		return
	}
	now2 := time.Now().Unix()

	// update changes to db
	_, err2 := database.DB.Exec(queries.ResetPassword(u.ID, hashed, now2))
	if err2 != nil {
		abort.Err(400, err2, ctx)
		return
	}

	createSendToken(u, ctx)
}

// auth protection middleware
func Protect(ctx *fiber.Ctx) {
	// 1) check if there is a token and it is valid
	token := ctx.Cookies("jwt", "empty")
	if token == "empty" {
		abort.Msg(401, "You are not logged in! Please log in to get access.", ctx)
	}
	// token Verification
	valid, tokenDetails := jwt.VerifyToken(token)
	if !valid {
		abort.Msg(401, "invalid token.", ctx)
	}

	// 2) Check if user still exists
	id := int(tokenDetails.Id.(float64))
	iat := tokenDetails.Iat.(float64)

	result := user{}
	err := database.DB.Get(&result, queries.GetUserWithId(id))
	if err != nil {
		abort.Msg(401, "The user belonging to this token does no longer exist.", ctx)
		return
	}
	// 3) Check if user changed password after the token was issued
	if password.ChangedPasswordAfter(iat, result.Password_changed_at) {
		abort.Msg(401, "User recently changed password! Please log in again.", ctx)
	}

	// GRANT ACCESS TO PROTECTED ROUTE
	ctx.Locals("user", result)
	ctx.Next()
}
