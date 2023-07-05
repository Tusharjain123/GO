package controller

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		panic("Error loading .env file")
	}
	return os.Getenv(key)
}

const dbName = "user"
const collectionName = "LoginCreds"

var collection *mongo.Collection

// Only run when initialisation
func init() {
	connectionString := goDotEnvVariable("MONGO_URL")
	// This step does not fire up a connection request
	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	// context.TODO => when you ae unclear what operation you are performing
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected Successfully")

	collection = client.Database(dbName).Collection(collectionName)

	fmt.Println("collection reference is ready")
}
