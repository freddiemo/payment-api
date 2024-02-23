package service

import (
	"errors"

	"github.com/freddiemo/payment-api/internal/transactions/model"
	"github.com/freddiemo/payment-api/internal/transactions/repository"
	"github.com/go-faker/faker/v4"
)

type TransactionsService interface {
	Payment(*model.Transaction) error
	FindPaymentById(id uint) (*model.Transaction, error)
	Refund(platformTransactionID string) (*model.Transaction, error)
}

type transactionsServ struct {
	transactionsRepository repository.TransactionsRepository
}

func NewTransactionsService(
	repository repository.TransactionsRepository,
) TransactionsService {
	return &transactionsServ{
		transactionsRepository: repository,
	}
}

func (service *transactionsServ) Payment(transaction *model.Transaction) error {
	// TODO: bank integration
	transaction.Type = "payment"
	transaction.Status = "successfull"
	transaction.PaymentPlatformTransactionID = faker.CCNumber()

	if err := service.transactionsRepository.Save(transaction); err != nil {
		return err
	}

	return nil
}

func (service *transactionsServ) FindPaymentById(id uint) (*model.Transaction, error) {
	var transaction *model.Transaction
	transaction, err := service.transactionsRepository.FindPaymentById(id)
	if err != nil {
		return transaction, err
	}

	return transaction, nil
}

func (service *transactionsServ) Refund(platformTransactionID string) (*model.Transaction, error) {
	var transaction *model.Transaction
	transaction, err := service.transactionsRepository.FindLastTransactionByPlatformTransactionID(platformTransactionID)
	if err != nil {
		return nil, err
	}

	if transaction.Type == "refund" {
		return nil, errors.New("Transaction has already refunded.")
	}

	// TODO: bank integration
	transaction.Type = "refund"
	transaction.Status = "successfull"

	if err := service.transactionsRepository.Save(transaction); err != nil {
		return transaction, err
	}

	return transaction, nil
}
