package entity

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID          string `gorm:"type:varchar(36)" json:"id"`
	Title       string `gorm:"type:varchar(255)" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	UserID      string `gorm:"not null" json:"-"`
	User        User   `gorm:"foreignKey:UserId;constraint:onUpdate:CASCADE,onDelete:CASCADE" json:"user"`
}

func (book *Book) BeforeCreate(db *gorm.DB) (err error) {
	book.ID = uuid.New().String()
	book.ID = uuid.New().String()
	return
}
