package mongodb

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	databaseUser     = os.Getenv("MONGODB_USER")
	databasePassword = os.Getenv("MONGODB_PASSWORD")
	databaseUri      = "mongodb://" + databaseUser + ":" + databasePassword + "@" + os.Getenv("MONGODB_URI")
)

func Connect() (*mongo.Client, context.Context, error) {
	log.Print("Mongodb client connecting...")
	client, err := mongo.NewClient(options.Client().ApplyURI(databaseUri))
	if err != nil {
		log.Fatal(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	if cancel != nil {
		log.Fatal(cancel)
	}

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return client, ctx, err
}
