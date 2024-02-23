package controller

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/freddiemo/payment-api/internal/transactions/model"
	"github.com/freddiemo/payment-api/internal/transactions/service"
)

type TransactionsController interface {
	Payment(ctx *gin.Context) (model.Transaction, error)
	PaymentDetail(ctx *gin.Context) (*model.Transaction, error)
	Refund(ctx *gin.Context) (*model.Transaction, error)
}

type transactionsController struct {
	transactionsService service.TransactionsService
}

func NewTransactionsController(transactionsService service.TransactionsService) TransactionsController {
	return &transactionsController{
		transactionsService: transactionsService,
	}
}

func (controller *transactionsController) Payment(ctx *gin.Context) (model.Transaction, error) {
	var transaction model.Transaction
	if err := ctx.ShouldBindJSON(&transaction); err != nil {
		return model.Transaction{}, err
	}

	if err := controller.transactionsService.Payment(&transaction); err != nil {
		return model.Transaction{}, err
	}

	return transaction, nil
}

func (controller *transactionsController) PaymentDetail(ctx *gin.Context) (*model.Transaction, error) {
	var transaction *model.Transaction

	id, err := strconv.ParseInt(ctx.Param("id"), 0, 0)
	if err != nil {
		return &model.Transaction{}, err
	}

	transaction, err = controller.transactionsService.FindPaymentById(uint(id))
	if err != nil {
		return &model.Transaction{}, err
	}

	return transaction, nil
}

func (controller *transactionsController) Refund(ctx *gin.Context) (*model.Transaction, error) {
	var refund model.Refund
	if err := ctx.ShouldBindJSON(&refund); err != nil {
		return nil, err
	}

	var transaction *model.Transaction
	transaction, err := controller.transactionsService.Refund(refund.PaymentPlatformTransactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}
