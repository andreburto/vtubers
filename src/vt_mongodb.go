package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/gorilla/mux"
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

type Generation2 struct {
	Id        int
	Name      string
	Company   string
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

func GetRoot2(w http.ResponseWriter, r *http.Request) {
	var html string = `<h1>VTubers</h1>
	<p><a href="/v2/company">Companies</a></p>
	<p><a href="/v2/generation">Generations</a></p>
	<p><a href="/v2/vtuber">VTubers</a></p>`
	DisplayPage(w, MakeHtml(html))
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

func CompanyIdMongoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	if _, ok := vars["id"]; !ok {
		http.Error(w, "Company ID not provided", http.StatusBadRequest)	
		return
	}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid Company ID", http.StatusBadRequest)
		return
	}

	client := CreateMongoDBClient()

	defer DisconnectMongoDBClient(client)

	collection := client.Database("vtubers").Collection("companies")
	filter := bson.D{bson.E{Key: "Id", Value: id}}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		http.Error(w, "Error fetching company data", http.StatusInternalServerError)
		return
	}

	defer cursor.Close(context.TODO())

	var company Company2
	err2 := cursor.Decode(&company)
	if err2 != nil {
		log.Fatal(err2)
		http.Error(w, "Company not found", http.StatusNotFound)
		return
	}

	var template string = `<h1>Company</h1>
	<h2>%s</h2>
	<hr>
	<p><a href="/v2/company">Back to Companies</a></p>
	<p><a href="/v2/">V2 Home</a></p>`
	
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, company.Name)))
}

func GenerationMongoHandler(w http.ResponseWriter, r *http.Request) {
	client := CreateMongoDBClient()

	defer DisconnectMongoDBClient(client)

	collection := client.Database("vtubers").Collection("generations")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var generationList string = ""
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var generation Generation2
		err := cursor.Decode(&generation)
		if err != nil {
			log.Fatal(err)
		}
		generationList = fmt.Sprintf("%s<p><a href=\"/generation/%d\">%s</a></p>", generationList, generation.Id, generation.Name)
	}

	var template string = `<h1>Generations</h1>
	%s
	<hr>
	<p><a href="/v2/">V2 Home</a></p>`
	
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, generationList)))
}

func VTuberMongoHandler(w http.ResponseWriter, r *http.Request) {
	client := CreateMongoDBClient()

	defer DisconnectMongoDBClient(client)

	collection := client.Database("vtubers").Collection("vtubers")
	filter := bson.D{}

	cursor, err := collection.Find(context.TODO(), filter)

	if err != nil {
		log.Fatal(err)
		panic(err)
	}

	var vtuberList string = ""
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var vtuber VTuber2
		err := cursor.Decode(&vtuber)
		if err != nil {
			log.Fatal(err)
		}
		vtuberList = fmt.Sprintf("%s<p><a href=\"/vtuber/%d\">%s</a> (%s, %s)</p>", vtuberList, vtuber.Id, vtuber.Name, vtuber.Company, vtuber.Generation)
	}

	var template string = `<h1>VTubers</h1>
	%s
	<hr>
	<p><a href="/v2/">V2 Home</a></p>`
	
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, vtuberList)))
}
