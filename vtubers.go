package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	LoadData()

	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot)
	router.HandleFunc("/company", CompanyHandler)
	// router.HandleFunc("/company/add", CompanyAddHandler)
	router.HandleFunc("/company/{id}", CompanyIdHandler)
	router.HandleFunc("/company/{id}/vtuber", CompanyIdVTuberHandler)
	router.HandleFunc("/company/{id}/generation", CompanyIdGenerationHandler)
	router.HandleFunc("/generation", GenerationHandler)
	router.HandleFunc("/generation/{id}", GenerationIdHandler)
	router.HandleFunc("/vtuber", VTuberHandler)
	router.HandleFunc("/vtuber/{id}", VTuberIdHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "0.0.0.0:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
