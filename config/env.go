package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURI() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	var mongoUri = os.Getenv("MONGOURI")

	log.Println(mongoUri)

	return mongoUri
}

func EnvDBName() string {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env file")
	}

	var dbNamee = os.Getenv("DB")

	log.Println(dbNamee)

	return dbNamee
}
