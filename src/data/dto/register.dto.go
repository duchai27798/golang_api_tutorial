package dto

type RegisterDTO struct {
	Name     string `json:"name" from:"name" validate:"required,gte=4"`
	Email    string `json:"email" from:"email" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"required,gte=6"`
}
