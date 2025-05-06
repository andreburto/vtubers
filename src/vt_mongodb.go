package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	// "go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CompanyMongoHandler(w http.ResponseWriter, r *http.Request) {
	var template string = `<h1>Companies</h1>
	<p><a href="/company/add">Add Company</a></p>
	%s
	<hr>
	<p><a href="/v2/">V2 Home</a></p>`
	var companyList string = ""
	// for _, company := range companies {
	// 	companyList = fmt.Sprintf("%s<p><a href="/company/%d">%s</a></p>", companyList, company.Id, company.Name)
	// }
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

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	DisplayPage(w, "Connected to MongoDB!")
}
