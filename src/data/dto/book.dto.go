package dto

import "github.com/google/uuid"

type BookUpdateDTO struct {
	ID          uuid.UUID `json:"id" from:"id" binding:"required"`
	Title       string    `json:"title" from:"title" binding:"required"`
	Description string    `json:"description" from:"description" binding:"required"`
	UserID      uuid.UUID `json:"user_id,omitempty" from:"user_id,omitempty"`
}

type BookCreateDTO struct {
	Title       string    `json:"title" from:"title" binding:"required"`
	Description string    `json:"description" from:"description" binding:"required"`
	UserID      uuid.UUID `json:"user_id,omitempty" from:"user_id,omitempty"`
}
