package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongo() {
	// Setup the mgm default config
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	// fmt.Println("uri:", os.Getenv("MONGOURI"))
	errMongo := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGOURI")))

	if errMongo != nil {
		log.Fatal("Error connecting to mongo")
	} else {
		fmt.Println("Connected to db")
	}
}
