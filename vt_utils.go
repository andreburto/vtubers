package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type Company struct {
	Id		int		`csv:"Id"`
	Name	string	`csv:"Name"`
}

type Generation struct {
	Id			int		`csv:"Id"`
	Name		string	`csv:"Name"`
	CompanyId	int		`csv:"CompanyId"`
}

type VTuber struct {
	Id				int		`csv:"Id"`
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
