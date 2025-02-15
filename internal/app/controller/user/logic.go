package user

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lutracorp/aonyx/internal/pkg/database"
)

// GetCurrent returns the user object of the requester's account.
func (c *Controller) GetCurrent(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*database.User)

	return ctx.JSON(user)
}

// ModifyCurrent modifies the requester's user account.
func (c *Controller) ModifyCurrent(ctx *fiber.Ctx) error {
	user := ctx.Locals("user").(*database.User)

	body := ModifyBody{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := c.Validator.Validate(&body); err != nil {
		return err
	}

	if body.Name != "" {
		user.Name = body.Name
	}

	if body.Email != "" {
		user.Email = body.Email
	}

	if body.Password != "" {
		encoded, err := c.Argon.HashEncoded([]byte(body.Password))
		if err != nil {
			return err
		}

		user.PasswordHash = string(encoded)
	}

	if tx := database.DB.Save(&user); tx.Error != nil {
		return tx.Error
	}

	return ctx.JSON(user)
}
