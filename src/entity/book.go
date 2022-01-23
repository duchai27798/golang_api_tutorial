package entity

import "github.com/google/uuid"

type Book struct {
	ID          uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4()" json:"id"`
	Title       string    `gorm:"type:varchar(255)" json:"title"`
	Description string    `gorm:"type:text" json:"description"`
	UserID      uuid.UUID `gorm:"not null" json:"-"`
	User        User      `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE" json:"user"`
}
