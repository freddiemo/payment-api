package model

import (
	"time"

	"gorm.io/gorm"
)

type PaymentPlatform struct {
	ID          uint           `gorm:"primaryKey"`
	CreatedAt   time.Time      `json:"-"`
	UpdatedAt   time.Time      `json:"-"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	PlatformKey string         `binding:"required" gorm:"not null"`
	Type        string         `binding:"required" gorm:"not null"`
	MerchantID  uint           `binding:"required" gorm:"not null" json:"merchantID"`
}
