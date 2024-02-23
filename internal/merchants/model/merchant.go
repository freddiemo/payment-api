package model

import (
	"time"

	"gorm.io/gorm"
)

type Merchant struct {
	ID         uint           `gorm:"primaryKey"`
	CreatedAt  time.Time      `json:"-"`
	UpdatedAt  time.Time      `json:"-"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	Name       string         `gorm:"type:varchar(32);not null"`
	Email      string         `gorm:"type:varchar(32);not null;UNIQUE"`
	DocumentID string         `gorm:"not null;UNIQUE"`
	Phone      string         `gorm:"not null"`
}
