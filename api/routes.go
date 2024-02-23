package api

import (
	"github.com/gin-gonic/gin"
)

func addPaymentRoutes(rg *gin.RouterGroup, paymentAPI *PaymentAPI) {
	transactions := rg.Group("/transactions")
	{
		transactions.POST("/payments", paymentAPI.Payment)
		transactions.GET("/payments/:id", paymentAPI.PaymentDetail)

		transactions.POST("/refunds", paymentAPI.Refund)
	}
}

func getRegisterRoutes(paymentAPI *PaymentAPI, server *gin.Engine) {
	v1 := server.Group("/v1")
	addPaymentRoutes(v1, paymentAPI)
}
