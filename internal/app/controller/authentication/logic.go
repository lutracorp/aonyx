package authentication

import (
	"github.com/gofiber/fiber/v2"
	"github.com/lutracorp/aonyx/internal/pkg/database"
	"github.com/lutracorp/aonyx/pkg/token"
	"github.com/lutracorp/foxid-go"
	"github.com/matthewhartstonge/argon2"
)

// Register creates a new account.
func (c *Controller) Register(ctx *fiber.Ctx) error {
	body := RegisterBody{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := c.Validator.Validate(body); err != nil {
		return err
	}

	id := foxid.Generate()

	encoded, err := c.Argon.HashEncoded([]byte(body.Password))
	if err != nil {
		return err
	}

	user := &database.User{
		ID:           id.String(),
		Name:         body.Name,
		Email:        body.Email,
		PasswordHash: string(encoded),
	}

	if tx := database.DB.Create(user); tx.Error != nil {
		return tx.Error
	}

	tok, err := token.Sign(id.Bytes(), encoded)
	if err != nil {
		return err
	}

	return ctx.JSON(&TokenResponse{tok})
}

// Login retrieves an authentication token for the given credentials.
func (c *Controller) Login(ctx *fiber.Ctx) error {
	body := LoginBody{}
	if err := ctx.BodyParser(&body); err != nil {
		return err
	}
	if err := c.Validator.Validate(body); err != nil {
		return err
	}

	user := &database.User{}
	if tx := database.DB.Where("email = ?", body.Email).First(user); tx.Error != nil {
		return tx.Error
	}

	ok, err := argon2.VerifyEncoded([]byte(body.Password), []byte(user.PasswordHash))
	if err != nil {
		return err
	}

	if !ok {
		return ctx.SendStatus(fiber.StatusUnauthorized)
	}

	id, err := foxid.Parse(user.ID)
	if err != nil {
		return err
	}

	tok, err := token.Sign(id.Bytes(), []byte(user.PasswordHash))
	if err != nil {
		return err
	}

	return ctx.JSON(&TokenResponse{tok})
}
