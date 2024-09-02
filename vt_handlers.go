package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

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