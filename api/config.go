package api

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func loadEnvVariables() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalln("Error loading .env file:", err)
	}
}

func envACCOUNTSID() string {
	loadEnvVariables()
	return os.Getenv("TWILIO_ACCOUNT_SID")
}


func envAUTHTOKEN() string {
	loadEnvVariables()
	return os.Getenv("TWILIO_AUTHTOKEN")
}


func envSERVICESID() string {
	loadEnvVariables()
	return os.Getenv("TWILIO_SERVICES_ID")
}