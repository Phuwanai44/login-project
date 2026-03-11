package config

import (
	"context"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database
var UserCollection *mongo.Collection

func ConnectDB() {

	uri := os.Getenv("MONGO_URI")

	client, err := mongo.Connect(
		context.TODO(),
		options.Client().ApplyURI(uri),
	)

	if err != nil {
		panic(err)
	}

	DB = client.Database("authDB")
	UserCollection = DB.Collection("users")

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		panic(err)
	}

	fmt.Println("MongoDB Connected")
	fmt.Println("Mongo URI:", os.Getenv("MONGO_URI"))
}
