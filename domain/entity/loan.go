package entity

import (
	"time"

	"gorm.io/gorm"
)

type Loan struct {
	ID        uint   `gorm:"not null;uniqueIndex;primaryKey" json:"id"`
	BookID    uint   `gorm:"not null;" json:"book_id"`
	UserID    uint   `gorm:"not null;" json:"user_id"`
	StartDate string `gorm:"not null;" json:"start_date"`
	EndDate   string `gorm:"not null;" json:"end_date"`
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `json:"deleted_at"`
}
