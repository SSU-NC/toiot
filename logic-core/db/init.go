package db

import (
	"context"
	"fmt"
	"reflect"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var cli *mongo.Client

func Setup() {
		// Set client options
		clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

		// Connect to MongoDB
		cli, err := mongo.Connect(context.TODO(), clientOptions)
		fmt.Println("\nresult type:", reflect.TypeOf(cli))
		if err != nil {
			fmt.Println("connect error: ", err.Error())
		}
	
		// Check the connection
		err = cli.Ping(context.TODO(), nil)
	
		if err != nil {
			fmt.Println("connect error: ", err.Error())
		}
	
		fmt.Println("Connected to MongoDB!")
		fmt.Println("cli:", cli)
}