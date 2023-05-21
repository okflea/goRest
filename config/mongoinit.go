package config

import (
	"fmt"
	"log"
	"os"
	"github.com/joho/godotenv"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/bson"
)

func InitMongo() {
	// Setup the mgm default config
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Fatal("Error loading .env file")
	}
	errMongo := mgm.SetDefaultConfig(nil, "mgm_lab", options.Client().ApplyURI(os.Getenv("MONGOURI")))

	if errMongo != nil {
		log.Fatal("Error connecting to mongo")
	} else {
		fmt.Println("Connected to db")
	}
}


type User struct {
	// DefaultModel adds _id, created_at and updated_at fields to the Model.
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name" bson:"name"`
 }
 
 func NewUser(name string) *User {
	return &User{
	   Name: name,
	}
 }

func CreateUser(name string) {
	user := NewUser(name)
	fmt.Println("user",user)
	err := mgm.Coll(user).Create(user)
	fmt.Println("err",err)
}

func FindUser(){
	user := &User{}
	coll := mgm.Coll(user)

	errResponse := coll.First(bson.M{}, user)

	if errResponse != nil {
		fmt.Println(errResponse)
	} else {
		fmt.Println(user)
	}
}
