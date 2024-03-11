package conn

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

func MongoConnect() error {
	clientOptions := options.Client().ApplyURI("mongodb://mongo-root:mongo-password@localhost:27017")
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return err
	}

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return err
	}

	MongoClient = client
	log.Println("Connected to MongoDB")
	return nil
}
