package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	transactions "github.com/freddiemo/payment-api/internal/transactions/controller"
)

type PaymentAPI struct {
	transactionsController transactions.TransactionsController
}

func NewPaymentAPI(
	transactionsController transactions.TransactionsController,
) *PaymentAPI {
	return &PaymentAPI{
		transactionsController: transactionsController,
	}
}

func (api *PaymentAPI) Payment(ctx *gin.Context) {
	transaction, err := api.transactionsController.Payment(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, transaction)
	}
}

func (api *PaymentAPI) PaymentDetail(ctx *gin.Context) {
	transaction, err := api.transactionsController.PaymentDetail(ctx)
	if err != nil {
		ctx.JSON(http.StatusNotFound, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, transaction)
	}
}

func (api *PaymentAPI) Refund(ctx *gin.Context) {
	transaction, err := api.transactionsController.Refund(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, &Response{
			Message: err.Error(),
		})
	} else {
		ctx.JSON(http.StatusCreated, transaction)
	}
}
