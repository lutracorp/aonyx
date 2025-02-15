package middleware

import (
	"github.com/gofiber/fiber/v2"
	tokenpb "github.com/lutracorp/aonyx/api/protocol/pkg/token"
	"github.com/lutracorp/aonyx/internal/pkg/database"
	"github.com/lutracorp/aonyx/pkg/token"
	"github.com/lutracorp/foxid-go"
)

func User(ctx *fiber.Ctx) error {
	tok := ctx.Get(fiber.HeaderAuthorization)

	data := &tokenpb.Data{}
	if err := token.Unmarshal(tok, data); err != nil {
		return err
	}

	if len(data.Payload) != 16 {
		return fiber.ErrUnauthorized
	}

	rid := foxid.FOxID(data.Payload)
	sid := rid.String()

	user := &database.User{}
	if tx := database.DB.Where("id = ?", sid).First(user); tx.Error != nil {
		return tx.Error
	}

	if ok, _ := token.VerifyData(data, []byte(user.PasswordHash)); !ok {
		return fiber.ErrUnauthorized
	}

	ctx.Locals("user", user)

	return ctx.Next()
}
