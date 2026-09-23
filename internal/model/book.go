package model

import (
	"time"
)

type Book struct {
	ID          int       `gorm:"primaryKey" json:"id"`
	Title       string    `json:"title" binding:"required"`
	Price       int       `json:"price" binding:"required,gt=0"`
	PageCount   int       `json:"page_count" binding:"required,gt=0"`
	Description string    `json:"description"`
	Rating      int       `json:"rating"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
