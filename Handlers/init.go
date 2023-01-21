package Handlers

import (
	"MongoDb/pkg/logging"
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var client *mongo.Client
var users *mongo.Collection

func Init(dbName, collectionName string) {
	var err error
	logger := logging.GetLogger()
	client, err = mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	users = client.Database(dbName).Collection(collectionName)
	logger.Infof("Connected to a database: <%s>, collection: <%s>", dbName, collectionName)
}
