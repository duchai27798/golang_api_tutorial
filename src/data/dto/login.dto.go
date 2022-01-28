package dto

type LoginDTO struct {
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:6" binding:"required"`
}
