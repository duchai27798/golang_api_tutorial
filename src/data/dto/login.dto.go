package dto

type LoginDTO struct {
	Email    string `json:"email" from:"email" validate:"required,email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"required,gte=6"`
}
