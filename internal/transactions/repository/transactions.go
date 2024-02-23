package repository

import (
	"gorm.io/gorm"

	"github.com/freddiemo/payment-api/internal/transactions/model"
)

type TransactionsRepository interface {
	Save(*model.Transaction) error
	FindPaymentById(uint) (*model.Transaction, error)
	FindLastTransactionByPlatformTransactionID(platformTransactionID string) (*model.Transaction, error)
}

type transactionsRepo struct {
	db *gorm.DB
}

func NewTransactionsRepository(db *gorm.DB) TransactionsRepository {
	return &transactionsRepo{
		db: db,
	}
}

func (transactionsRepo *transactionsRepo) Save(transaction *model.Transaction) error {
	if result := transactionsRepo.db.Save(transaction); result.Error != nil {
		return result.Error
	}

	return nil
}

func (transactionsRepo *transactionsRepo) FindPaymentById(id uint) (*model.Transaction, error) {
	var transaction model.Transaction
	if result := transactionsRepo.db.Where("type = ?", "payment").First(&transaction, id); result.Error != nil {
		return &transaction, result.Error
	}

	return &transaction, nil
}

func (transactionsRepo *transactionsRepo) FindLastTransactionByPlatformTransactionID(platformTransactionID string) (*model.Transaction, error) {
	var transaction model.Transaction
	if result := transactionsRepo.db.Where("payment_platform_transaction_id = ?", platformTransactionID).First(&transaction); result.Error != nil {
		return &transaction, result.Error
	}

	return &transaction, nil
}
