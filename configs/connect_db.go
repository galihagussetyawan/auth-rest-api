package configs

import (
	"context"
	"log"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Client = ConnectDB()

func ConnectDB() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Not found file .env")
	}

	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set yout 'MONGODB_URI' env")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	defer func() {
		err := client.Disconnect(context.TODO())
		if err != nil {
			panic(err)
		}
	}()

	return client
}
