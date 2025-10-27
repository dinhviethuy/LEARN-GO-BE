package model

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey;autoIncrement" json:"id"`
	Name      string         `gorm:"size:100;not null" json:"name" validate:"required,min=2"`
	Email     string         `gorm:"size:100;not null" json:"email" validate:"required,email"`
	Password  string         `gorm:"not null" json:"-" validate:"required,min=6"`
	Age       int            `json:"age" validate:"gte=0,lte=120"`
	Posts     []Post         `gorm:"foreignKey:UserID" json:"posts,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

