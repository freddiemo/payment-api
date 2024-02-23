package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

var envFile = ".env"

func Init() map[string]string {
	params, err := godotenv.Read(envFile)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading %s file", envFile))
	}

	if os.Getenv("GO_ENVIRONMENT") == "TEST" {
		params = SetTestParams(params)
	}

	return params
}

func SetTestParams(params map[string]string) map[string]string {
	sufix := "_TEST"
	params["DB_HOST"] = params[fmt.Sprintf("DB_HOST%s", sufix)]
	params["DB_USER"] = params[fmt.Sprintf("DB_USER%s", sufix)]
	params["DB_PASSWORD"] = params[fmt.Sprintf("DB_PASSWORD%s", sufix)]
	params["DB_NAME"] = params[fmt.Sprintf("DB_NAME%s", sufix)]
	params["DB_PORT"] = params[fmt.Sprintf("DB_PORT%s", sufix)]

	return params
}
