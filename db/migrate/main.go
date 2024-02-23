package main

import (
	"log"

	"github.com/freddiemo/payment-api/config"
	"github.com/freddiemo/payment-api/db"

	merchantsSeeder "github.com/freddiemo/payment-api/db/seeders"
	merchants "github.com/freddiemo/payment-api/internal/merchants/model"
	paymentPlatforms "github.com/freddiemo/payment-api/internal/payment-platforms/model"

	customersSeeder "github.com/freddiemo/payment-api/db/seeders"
	customers "github.com/freddiemo/payment-api/internal/customers/model"

	transactions "github.com/freddiemo/payment-api/internal/transactions/model"
)

func main() {
	params := config.Init()
	db := db.Init(params)

	db.AutoMigrate(&merchants.Merchant{})
	db.AutoMigrate(&paymentPlatforms.PaymentPlatform{})
	var countMerchants int64
	if result := db.Model(&merchants.Merchant{}).Count(&countMerchants); result.Error != nil {
		log.Fatal(result.Error)
	} else if countMerchants == 0 {
		merchantSeeder := merchantsSeeder.NewMerchantsSeeder(db)
		if err := merchantSeeder.Populate(); err != nil {
			log.Fatal(err)
		}
	}

	db.AutoMigrate(&customers.Customer{})
	var countCustomers int64
	if result := db.Model(&customers.Customer{}).Count(&countCustomers); result.Error != nil {
		log.Fatal(result.Error)
	} else if countCustomers == 0 {
		customerSeeder := customersSeeder.NewCustomersSeeder(db)
		if err := customerSeeder.Populate(); err != nil {
			log.Fatal(err)
		}
	}

	db.AutoMigrate(&transactions.Transaction{})

	log.Println("Successfully migrated.")
}
