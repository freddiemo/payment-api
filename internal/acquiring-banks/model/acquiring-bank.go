package model

import (
	"time"

	"gorm.io/gorm"
)

/*type BankCompany string

const (
	Bancolombia BankCompany = "bancolombia"
	BancoBogota BankCompany = "banco-bogota"
	BBVA        BankCompany = "bbva"
	Citibank    BankCompany = "citibank"
)*/

type AcquiringBank struct {
	ID            uint           `gorm:"primaryKey"`
	CreatedAt     time.Time      `json:"-"`
	UpdatedAt     time.Time      `json:"-"`
	DeletedAt     gorm.DeletedAt `gorm:"index" json:"-"`
	DateTime      *time.Time     `binding:"required" gorm:"not null" json:"dateTime"`
	MerchantID    uint           `binding:"required" json:"merchantID"`
	CustomerID    uint           `binding:"required" json:"customerID"`
	NumberAccount uint           `binding:"required" gorm:"not null"`
}
