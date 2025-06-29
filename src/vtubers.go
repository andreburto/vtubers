package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadData()

	router := mux.NewRouter()
	// v1, csv version
	router.HandleFunc("/", GetRoot)
	router.HandleFunc("/company", CompanyHandler)
	router.HandleFunc("/company/add", CompanyAddHandler)
	router.HandleFunc("/company/{id}", CompanyIdHandler)
	// router.HandleFunc("/company/{id}/edit", CompanyIdEditHandler)
	router.HandleFunc("/company/{id}/vtuber", CompanyIdVTuberHandler)
	router.HandleFunc("/company/{id}/generation", CompanyIdGenerationHandler)
	router.HandleFunc("/generation", GenerationHandler)
	router.HandleFunc("/generation/add", GenerationAddHandler)
	router.HandleFunc("/generation/{id}", GenerationIdHandler)
	router.HandleFunc("/vtuber", VTuberHandler)
	router.HandleFunc("/vtuber/add", VTuberAddHandler)
	router.HandleFunc("/vtuber/{id}", VTuberIdHandler)
	// v2, mongo version
	router.HandleFunc("/v2/", GetRoot2)
	router.HandleFunc("/v2/company", CompanyMongoHandler)
	router.HandleFunc("/v2/company/{id}", CompanyIdMongoHandler)
	router.HandleFunc("/v2/generation", GenerationMongoHandler)
	router.HandleFunc("/v2/vtuber", VTuberMongoHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
