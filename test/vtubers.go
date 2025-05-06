package main

import (
	"context"
	"fmt"
	"log"
	"os"
	// "time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VTuber struct {
	Id         int
	Name       string
	Company    string 
	Generation string
}

func main() {
	fmt.Println(os.Getenv("MONGO_URI"))
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err = client.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	collection := client.Database("vtubers").Collection("vtubers")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.TODO())

	for cursor.Next(context.TODO()) {
		var result VTuber
		err := cursor.Decode(&result)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s (%s)\n", result.Name, result.Company)
	}
}
