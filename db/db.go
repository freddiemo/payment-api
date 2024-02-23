package db

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(params map[string]string) *gorm.DB {
	dsn := GetDsn(params)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})

	if err != nil {
		log.Fatal("Failed to connect database")
	}

	return db
}

func GetDsn(params map[string]string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s", params["DB_HOST"], params["DB_USER"], params["DB_PASSWORD"], params["DB_NAME"], params["DB_PORT"])
}
