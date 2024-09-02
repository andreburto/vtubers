package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func GetRoot(w http.ResponseWriter, r *http.Request) {
	var html string = `<h1>VTubers</h1>
<p><a href="/company">Companies</a></p>
<p><a href="/generation">Generations</a></p>
<p><a href="/vtuber">VTubers</a></p>`
	DisplayPage(w, MakeHtml(html))
}

func CompanyHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>Companies</h1>%s<hr><p><a href=\"/company/add\">Add Company</a></p>"
	var companyList string = ""
	for _, company := range companies {
		companyList = fmt.Sprintf("%s<p><a href=\"/company/%d\">%s</a></p>", companyList, company.Id, company.Name)
	}
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, companyList)))
}

func CompanyAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Retrieve form data
		companyName := r.Form.Get("company_name")

		// Validate form data: company name is required.
		if len(companyName) == 0 {
			http.Error(w, "Company name is required", http.StatusBadRequest)
			return
		}

		// This is a small app. I can take shortcuts.
		companyId := GetCompanyMaxId() + 1

		// Add the new company to the list
		companies = append(companies, &Company{Id: companyId, Name: companyName})

		// Save the data
		SaveData()

		http.Redirect(w, r, "/company", http.StatusSeeOther)
	} else {
		var form string = `<h1>Add Company</h1>
<form method="post">
<label for="company_name">Name:</label>
<input type="text" id="company_name" name="company_name">
<br><br>
<input type="submit" value="Add">
</form>`

		DisplayPage(w, MakeHtml(form))
	}
}

func CompanyIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	companyName, companyId := GetCompany(vars["id"])
	var companyPage string = `<h1>%s</h1>
<p><a href="/company/%d/generation">Generations</a></p>
<p><a href="/company/%d/vtuber">VTubers</a></p>`
	DisplayPage(w, MakeHtml(fmt.Sprintf(companyPage, companyName, companyId, companyId)))
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

	DisplayPage(w, MakeHtml(fmt.Sprintf("%s%s", header, body)))
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

	DisplayPage(w, MakeHtml(fmt.Sprintf("%s%s", header, body)))
}

func GenerationHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>Generations</h1>%s<hr><p><a href=\"/generation/add\">Add Generation</a></p>"
	var generationList string = ""
	for _, g := range generations {
		generationList = fmt.Sprintf("%s<p><a href=\"/generation/%d\">%s</a></p>", generationList, g.Id, g.Name)
	}
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, generationList)))
}

func GenerationAddHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Failed to parse form data", http.StatusBadRequest)
			return
		}

		// Retrieve form data
		generationName := r.Form.Get("generation_name")
		companyId, err := strconv.Atoi(r.Form.Get("company_id"))

		if err != nil {
			http.Error(w, "Failed to parse company id", http.StatusBadRequest)
			return
		}

		// Validate form data: generation name is required.
		if len(generationName) == 0 {
			http.Error(w, "Generation name is required", http.StatusBadRequest)
			return
		}

		// This is a small app. I can take shortcuts.
		generationId := GetGenerationMaxId() + 1

		// Add the new generation to the list
		generations = append(generations, &Generation{
			Id: generationId, 
			Name: generationName, 
			CompanyId: companyId})

		// Save the data
		SaveData()

		http.Redirect(w, r, "/generation", http.StatusSeeOther)
	} else {
		var companyOptions string = fmt.Sprintf("<select id=\"company_id\" name=\"company_id\">%s</select>", MakeCompanyOptions())
		var form string = `<h1>Add Generation</h1>
		<form method="post">
		<label for="generation_name">Name:</label>
		<input type="text" id="generation_name" name="generation_name">
		<br><br>
		<label for="company_id">Company:</label>
		%s
		<br><br>
		<input type="submit" value="Add">
		</form>`
		
		DisplayPage(w, MakeHtml(fmt.Sprintf(form, companyOptions)))
	}
}

func GenerationIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var g Generation = GetGeneration(vars["id"])
	DisplayPage(w, MakeHtml(fmt.Sprintf("<h1>%s</h1>", g.Name)))
}

func VTuberHandler(w http.ResponseWriter, r *http.Request) {
	var template string = "<h1>VTubers</h1>%s"
	var vtuberList string = ""
	for _, v := range vtubers {
		vtuberList = fmt.Sprintf("%s<p><a href=\"/vtuber/%d\">%s</a></p>", vtuberList, v.Id, v.Name)
	}
	DisplayPage(w, MakeHtml(fmt.Sprintf(template, vtuberList)))
}

func VTuberIdHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var v VTuber = GetVTuber(vars["id"])
	DisplayPage(w, MakeHtml(fmt.Sprintf("<h1>%s</h1>", v.Name)))
}
