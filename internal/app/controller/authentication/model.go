package authentication

type (
	RegisterBody struct {
		Name     string `json:"name" validate:"required,username,gte=2,lte=32"`
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,lte=72"`
	}

	LoginBody struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required,gte=8,lte=72"`
	}

	TokenResponse struct {
		Token string `json:"token"`
	}
)
