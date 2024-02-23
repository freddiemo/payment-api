package main

import (
	"github.com/freddiemo/payment-api/api"
	"github.com/freddiemo/payment-api/config"
	"github.com/freddiemo/payment-api/db"
)

func main() {
	params := config.Init()
	db := db.Init(params)
	api.Init(params, db)
}
