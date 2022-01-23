package entity

import "github.com/google/uuid"

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Name     string    `gorm:"type:varchar(255)" json:"name"`
	Email    string    `gorm:"uniqueIndex;type:varchar(255)" json:"email"`
	Password string    `gorm:"->;<-;not null" json:"-"`
	Token    string    `gorm:"-" json:"token,omitempty"`
}
