package db

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Abhishekkumar2021/golang-backend/utils"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var Database *mongo.Database

func init() {
	// Load the environment variables
	utils.LoadEnv()

	// Connect to the database
	connectionString := string(os.Getenv("MONGODB_URI"))
	clientOptions := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal("Error connecting to the database")
	}
	Database = client.Database(utils.DB_NAME)

	// Try Ping the database
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal("Error pinging the database")
	}
	fmt.Println("Connected to the database successfully")
}
