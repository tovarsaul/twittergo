package domain

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

var MongoDBConnection = Connect()
var clientOptions = options.Client().ApplyURI("mongodb+srv://saul:dT9aGdYN8sDhT00s@twitter.vmccd.mongodb.net/?retryWrites=true&w=majority")

func Connect() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Connection successful")
	return client
}

func CheckConection() int {
	err := MongoDBConnection.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
