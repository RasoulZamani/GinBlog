package config

import (
	"log"

	"github.com/joho/godotenv"
)

func ReadEnv() map[string]string {
	var myEnv map[string]string
	myEnv, err := godotenv.Read()

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return myEnv

}
