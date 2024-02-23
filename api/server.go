package api

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	transactionsCtrl "github.com/freddiemo/payment-api/internal/transactions/controller"
	transactionsRepo "github.com/freddiemo/payment-api/internal/transactions/repository"
	transactionsServ "github.com/freddiemo/payment-api/internal/transactions/service"
)

func Init(params map[string]string, db *gorm.DB) {
	setupLogOutput(params["APP_NAME"])
	server := GetServer(db)

	server.GET("/health", func(c *gin.Context) {
		c.String(200, "Healthy")
	})

	server.Run(fmt.Sprintf(":%s", params["APP_PORT"]))
}

func setupLogOutput(app_name string) {
	file, err := os.Create(fmt.Sprintf("%s.log", app_name))
	if err != nil {
		log.Fatal("Cannot create log file")
	}
	gin.DefaultWriter = io.MultiWriter(file, os.Stdout)
}

func GetServer(db *gorm.DB) *gin.Engine {
	var transactionsRepository transactionsRepo.TransactionsRepository = transactionsRepo.NewTransactionsRepository(db)
	var transactionsService transactionsServ.TransactionsService = transactionsServ.NewTransactionsService(transactionsRepository)
	var transactionsController transactionsCtrl.TransactionsController = transactionsCtrl.NewTransactionsController(transactionsService)

	server := gin.Default()
	paymentAPI := NewPaymentAPI(
		transactionsController,
	)
	getRegisterRoutes(paymentAPI, server)

	return server
}
