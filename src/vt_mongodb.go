package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VTuber2 struct {
	Id         int
	Name       string
	Company    string 
	Generation string
}

type Company2 struct {
	Id   int
	Name string
}

func CreateMongoDBClient() (*mongo.Client) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
	return client
}

func DisconnectMongoDBClient(client *mongo.Client) {
	err := client.Disconnect(context.TODO())
	if err != nil {
		log.Fatal(err)
		panic(err)
	}
}

func CompanyMongoHandler(w http.ResponseWriter, r *http.Request) {
	
	client := CreateMongoDBClient()

	defer DisconnectMongoDBClient(client)

	collection := client.Database("vtubers").Collection("companies")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var companyList string = ""
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var company Company2
		err := cursor.Decode(&company)
		if err != nil {
			log.Fatal(err)
		}
		companyList = fmt.Sprintf("%s<p><a href=\"/company/%d\">%s</a></p>", companyList, company.Id, company.Name)
	}

	var template string = `<h1>Companies</h1>
	%s
	<hr>
	<p><a href="/v2/">V2 Home</a></p>`
	
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, companyList)))
}

func GetRoot2(w http.ResponseWriter, r *http.Request) {
	var html string = `<h1>VTubers</h1>
	<p><a href="/v2/company">Companies</a></p>
	<p><a href="/v2/generation">Generations</a></p>
	<p><a href="/v2/vtuber">VTubers</a></p>`
	DisplayPage(w, MakeHtml(html))
}

func TestMongo(w http.ResponseWriter, r *http.Request) {
	client := CreateMongoDBClient()

	defer DisconnectMongoDBClient(client)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err := client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DisplayPage(w, "Connected to MongoDB!")
}
