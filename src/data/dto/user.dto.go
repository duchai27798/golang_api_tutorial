package dto

import "github.com/google/uuid"

type UserUpdateDTO struct {
	ID       uuid.UUID `json:"id" from:"id" binding:"required"`
	Name     string    `json:"name" from:"name" binding:"required" validate:"required"`
	Email    string    `json:"email" from:"email" binding:"required" validate:"email"`
	Password string    `json:"password,omitempty" from:"password,omitempty" validate:"min:6" binding:"required"`
}

type UserCreateDTO struct {
	Name     string `json:"name" from:"name" binding:"required" validate:"required"`
	Email    string `json:"email" from:"email" binding:"required" validate:"email"`
	Password string `json:"password,omitempty" from:"password,omitempty" validate:"min:6" binding:"required"`
}
