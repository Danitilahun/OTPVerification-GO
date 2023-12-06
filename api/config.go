package api

import (
	"log"
	"os"
	"github.com/joho/godotenv"
)

func envACCOUNTID() string {
	err:= godotenv.Load();
	if err != nil{
		log.Fatalln("Error loading .env file")
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_ACCOUNT_SID")
}


func envAUTHTOKEN() string {
	err:= godotenv.Load();
	if err != nil{
		log.Fatalln("Error loading .env file")
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_AUTHTOKEN")
}


func envSERVICEID() string {
	err:= godotenv.Load();
	if err != nil{
		log.Fatalln("Error loading .env file")
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("TWILIO_SERVICES_ID")
}