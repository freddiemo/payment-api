package model

import (
	"time"

	"gorm.io/gorm"
)

/*type TransactionType string

const (
	Payment TransactionType = "payment"
	Refund  TransactionType = "refund"
)

type StatusType string

const (
	Successfull  StatusType = "successfull"
	Pending  StatusType = "pending"
	Rejected StatusType = "rejected"
)*/

type Transaction struct {
	ID                           uint           `gorm:"primaryKey"`
	CreatedAt                    time.Time      `json:"-"`
	UpdatedAt                    time.Time      `json:"-"`
	DeletedAt                    gorm.DeletedAt `gorm:"index" json:"-"`
	PaymentPlatformTransactionID string         `gorm:"not null"`
	Amount                       float64        `binding:"required" gorm:"not null"`
	CustomerID                   uint           `binding:"required" gorm:"not null" json:"customerID"`
	MerchantID                   uint           `binding:"required" gorm:"not null" json:"merchantID"`
	PaymentPlatformID            uint           `binding:"required" gorm:"not null" json:"paymentPlatformID"`
	/*Type                         TransactionType `gorm:"type:enum('payment', 'refund')"`
	Status                       StatusType      `gorm:"type:enum('successfull', 'pending', 'rejected')"`*/
	Type                    string `gorm:"not null"`
	Status                  string `gorm:"not null"`
	PaymentPlatformResponse string
}

type Refund struct {
	PaymentPlatformTransactionID string `binding:"required" json:"platformTransactionID"`
}
