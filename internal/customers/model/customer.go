package model

import (
	"time"

	"gorm.io/gorm"
)

type Customer struct {
	ID         uint           `gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	FirstName  string         `binding:"required,min=3" gorm:"not null"`
	LastName   string         `binding:"required,min=3" gorm:"not null"`
	Email      string         `gorm:"type:varchar(32);not null;UNIQUE"`
	DocumentID string         `gorm:"not null;UNIQUE"`
	Phone      string         `gorm:"not null"`
	MerchantID uint           `binding:"required" gorm:"not null" json:"merchantID"`
}
