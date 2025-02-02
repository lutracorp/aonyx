package authentication

import (
	"github.com/lutracorp/aonyx/internal/app/util/validator"
	"github.com/matthewhartstonge/argon2"
)

// Controller represents authentication controller.
type Controller struct {
	Argon     *argon2.Config
	Validator *validator.Validator
}

// NewController creates new instance of authentication Controller.
func NewController() *Controller {
	arc := argon2.MemoryConstrainedDefaults()

	return &Controller{
		Argon:     &arc,
		Validator: validator.NewValidator(),
	}
}
