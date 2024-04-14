package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
	"github.com/gorilla/mux"
)

type Company struct {
	Id 		int		`csv:"Id"`
	Name	string	`csv:"Name"`
}

type Generation struct {
	Id 			int		`csv:"Id"`
	Name		string	`csv:"Name"`
	CompanyId	int		`csv:"CompanyId"`
}

type VTuber struct {
	Id 				int		`csv:"Id"`
	Name			string	`csv:"Name"`
	CompanyId		int		`csv:"CompanyId"`
	GenerationId	int		`csv:"GenerationId"`
}

// Global lists. Load once, use anywhere.
var companies []*Company
var generations []*Generation
var vtubers []*VTuber

// This function is ugly, but it works. Loads data from the three CSV files.
func LoadData() {
	companies = []*Company{}
	generations = []*Generation{}
	vtubers = []*VTuber{}

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	cin, err := os.Open(fmt.Sprintf("%s/company.csv", path))
	if err != nil {
		panic(err)
	}
	defer cin.Close()
	if err := gocsv.UnmarshalFile(cin, &companies); err != nil {
		panic(err)
	}
	
	gin, err := os.Open(fmt.Sprintf("%s/generation.csv", path))
	if err != nil {
		panic(err)
	}
	defer gin.Close()
	if err := gocsv.UnmarshalFile(gin, &generations); err != nil {
		panic(err)
	}
	
	vin, err := os.Open(fmt.Sprintf("%s/vtuber.csv", path))
	if err != nil {
		panic(err)
	}
	defer vin.Close()
	if err := gocsv.UnmarshalFile(vin, &vtubers); err != nil {
		panic(err)
	}
}

func GetCompany(id string) (string, int) {
	var companyName string = ""
	var companyId int = -1
	for _, company := range companies {
		if id == fmt.Sprintf("%d", company.Id) {
			companyName = company.Name
			companyId = company.Id
			break
		}
	}
	if companyName == "" {
		panic("No Company matches.")
	}
	return companyName, companyId
}

func GetGeneration(id string) Generation {
	var found bool = false
	var gen Generation
	for _, g := range generations {
		if id == fmt.Sprintf("%d", g.CompanyId) {
			found = true
			gen = (*g)
		}
	}
	if !found {
		panic("No Generation matches.")
	}
	return gen
}

func GetGenerationsByCompany(id string) ([]Generation) {
	var gs []Generation
	for _, g := range generations {
		if id == fmt.Sprintf("%d", g.Id) {
			gs = append(gs, (*g))
		}
	}
	if len(gs) == 0 {
		panic("No company matches.")
	}
	return gs
}

func GetVTuber(id string) VTuber {
	var found bool = false
	var v VTuber
	for _, vt := range vtubers {
		if id == fmt.Sprintf("%d", vt.Id) {
			found = true
			v = (*vt)
			break
		}
	}
	if !found {
		panic("No VTuber matches.")
	}
	return v
}

func GetVTubersByCompany(id string) ([]VTuber) {
	var vts []VTuber
	for _, vt := range vtubers {
		if id == fmt.Sprintf("%d", vt.CompanyId) {
			vts = append(vts, (*vt))
		}
	}
	if len(vts) == 0 {
		panic("No company matches.")
	}
	return vts
}

func makeHtml(msg string) string {
	return fmt.Sprintf("<html><head><title>Server</title></head><body>%s<hr><p><a href=\"/\">Home</a></body></html>", msg)
}

func displayPage(w http.ResponseWriter, c string) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, c)
}

func GetRoot(w http.ResponseWriter, r *http.Request) {
	var html string = `<h1>VTubers</h1>
<p><a href="/company">Companies</a></p>
<p><a href="/generation">Generations</a></p>
<p><a href="/vtuber">VTubers</a></p>`
	displayPage(w, makeHtml(html))
}

func CompanyHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>Companies</h1>%s"
	var companyList string = ""
	for _, company := range companies {
		companyList = fmt.Sprintf("%s<p><a href=\"/company/%d\">%s</a></p>", companyList, company.Id, company.Name)
	}
	displayPage(w, makeHtml(fmt.Sprintf(template, companyList)))
}

func CompanyIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName, companyId := GetCompany(vars["id"])
	var companyPage string = `<h1>%s</h1>
<p><a href="/company/%d/generation">Generations</a></p>
<p><a href="/company/%d/vtuber">VTubers</a></p>`
	displayPage(w, makeHtml(fmt.Sprintf(companyPage, companyName, companyId, companyId)))
}

func CompanyIdVTuberHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName, _ := GetCompany(vars["id"])
	var vts []VTuber = GetVTubersByCompany(vars["id"])
	var header string = fmt.Sprintf("<h1>%s</h1>", companyName)
	var body string = ""
	
	for _, vt := range vts {
		body = fmt.Sprintf("%s<p><a href=\"/vtuber/%d\">%s</a></p>", body, vt.Id, vt.Name) 
	}

	displayPage(w, makeHtml(fmt.Sprintf("%s%s", header, body)))
}

func CompanyIdGenerationHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName, _ := GetCompany(vars["id"])
	var gs []Generation = GetGenerationsByCompany(vars["id"])
	var header string = fmt.Sprintf("<h1>%s</h1>", companyName)
	var body string = ""
	
	for _, g := range gs {
		body = fmt.Sprintf("%s<p><a href=\"/generation/%d\">%s</a></p>", body, g.Id, g.Name) 
	}

	displayPage(w, makeHtml(fmt.Sprintf("%s%s", header, body)))
}

func GenerationHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>Generations</h1>%s"
	var generationList string = ""
	for _, g := range generations {
		generationList = fmt.Sprintf("%s<p><a href=\"/generation/%d\">%s</a></p>", generationList, g.Id, g.Name)
	}
	displayPage(w, makeHtml(fmt.Sprintf(template, generationList)))
}

func GenerationIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var g Generation = GetGeneration(vars["id"])
	displayPage(w, makeHtml(fmt.Sprintf("<h1>%s</h1>", g.Name)))
}

func VTuberHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>VTubers</h1>%s"
	var vtuberList string = ""
	for _, v := range vtubers {
		vtuberList = fmt.Sprintf("%s<p><a href=\"/vtuber/%d\">%s</a></p>", vtuberList, v.Id, v.Name)
	}
	displayPage(w, makeHtml(fmt.Sprintf(template, vtuberList)))
}

func VTuberIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var v VTuber = GetVTuber(vars["id"])
	displayPage(w, makeHtml(fmt.Sprintf("<h1>%s</h1>", v.Name)))
}

func main() {
	LoadData()

	router := mux.NewRouter()
	router.HandleFunc("/", GetRoot)
	router.HandleFunc("/company", CompanyHandler)
	router.HandleFunc("/company/{id}", CompanyIdHandler)
	router.HandleFunc("/company/{id}/vtuber", CompanyIdVTuberHandler)
	router.HandleFunc("/company/{id}/generation", CompanyIdGenerationHandler)
	router.HandleFunc("/generation", GenerationHandler)
	router.HandleFunc("/generation/{id}", GenerationIdHandler)
	router.HandleFunc("/vtuber", VTuberHandler)
	router.HandleFunc("/vtuber/{id}", VTuberIdHandler)

	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8080",
	}

	log.Fatal(srv.ListenAndServe())
}
