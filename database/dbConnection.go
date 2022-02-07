package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func DbInstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("error loading .env", err)
	}

	mongoDbUrl := os.Getenv("MONGO_DB")

	client, err := mongo.NewClient(options.Client().ApplyURI(mongoDbUrl))

	if err != nil {
		log.Fatal("unable to connect mongodb:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("unable to connect db", err)
	}
	fmt.Println("Connected to Mongodb:", mongoDbUrl)

	return client
}

var Client *mongo.Client = DbInstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var Collection *mongo.Collection = client.Database("cluster").Collection(collectionName)
	return Collection
}
