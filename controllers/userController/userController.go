package userController

import (
	"github.com/Badrouu17/go-postgresql-api-boilerplate/database"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/models"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/queries"
	"github.com/Badrouu17/go-postgresql-api-boilerplate/utils/abort"
	"github.com/gofiber/fiber"
)

func GetMe(ctx *fiber.Ctx) {
	u := ctx.Locals("user").(models.User)
	ctx.Status(200).JSON(&fiber.Map{
		"status": 200,
		"id":     u.ID,
		"name":   u.Name,
		"email":  u.Email,
	})
}

func UpdateMe(ctx *fiber.Ctx) {
	// getting the body the right way
	type updateMeInput struct {
		Name  string `json:"name"`
		Email string `json:"email"`
	}
	input := new(updateMeInput)
	ctx.BodyParser(input)
	// user from req Locales
	u := ctx.Locals("user").(models.User)
	// update name
	if input.Name != "" {
		_, err2 := database.DB.Exec(queries.UpdateUserName(u.ID, input.Name))
		if err2 != nil {
			abort.Err(400, err2, ctx)
			return
		}
	}
	// update email
	if input.Email != "" {
		_, err2 := database.DB.Exec(queries.UpdateUserEmail(u.ID, input.Email))
		if err2 != nil {
			abort.Err(400, err2, ctx)
			return
		}
	}
	ctx.Status(200).JSON(&fiber.Map{
		"status": 200,
		"msg":    "updated",
	})
}

func DeleteMe(ctx *fiber.Ctx) {
	u := ctx.Locals("user").(models.User)
	// delete from db
	_, err2 := database.DB.Exec(queries.DeleteUser(u.ID))
	if err2 != nil {
		abort.Err(400, err2, ctx)
		return
	}
	ctx.Status(200).JSON(&fiber.Map{
		"status": 200,
		"msg":    "deleted",
	})
}
