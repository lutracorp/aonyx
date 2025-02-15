package user

type ModifyBody struct {
	Name     string `json:"name" validate:"omitempty,username,gte=2,lte=32"`
	Email    string `json:"email" validate:"omitempty,email"`
	Password string `json:"password" validate:"omitempty,gte=8,lte=72"`
}
