package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gocarina/gocsv"
)

type Company struct {
	Id   int    `csv:"Id"`
	Name string `csv:"Name"`
}

type Generation struct {
	Id        int    `csv:"Id"`
	Name      string `csv:"Name"`
	CompanyId int    `csv:"CompanyId"`
}

type VTuber struct {
	Id           int    `csv:"Id"`
	Name         string `csv:"Name"`
	CompanyId    int    `csv:"CompanyId"`
	GenerationId int    `csv:"GenerationId"`
}

// Global lists. Load once, use anywhere.
var companies []*Company
var generations []*Generation
var vtubers []*VTuber

// Uncle Bob would be proud of this function.
func GetDataPath() string {
	var directory string = "data"

	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	return fmt.Sprintf("%s/%s", path, directory)
}

// This function is ugly, but it works. Loads data from the three CSV files.
func LoadData() {
	companies = []*Company{}
	generations = []*Generation{}
	vtubers = []*VTuber{}

	// Get the full path to the directory where my CSV files are kept.
	var path string = GetDataPath()

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

func SaveData() {
	// Get the full path to the directory where my CSV files are kept.
	var path string = GetDataPath()

	// Update the company CSV file.
	cout, err := os.Create(fmt.Sprintf("%s/company.csv", path))

	if err != nil {
		panic(err)
	}

	if err := gocsv.MarshalFile(&companies, cout); err != nil {
		panic(err)
	}

	cout.Close()

	// Update the generation CSV file.
	gout, err := os.Create(fmt.Sprintf("%s/generation.csv", path))

	if err != nil {
		panic(err)
	}

	if err := gocsv.MarshalFile(&generations, gout); err != nil {
		panic(err)
	}

	gout.Close()

	// Update the vtuber CSV file.
	vout, err := os.Create(fmt.Sprintf("%s/vtuber.csv", path))

	if err != nil {
		panic(err)
	}

	if err := gocsv.MarshalFile(&vtubers, vout); err != nil {
		panic(err)
	}

	vout.Close()
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

func GetCompanyMaxId() int {
	var maxId int = -1
	for _, company := range companies {
		if company.Id > maxId {
			maxId = company.Id
		}
	}
	return maxId
}

func GetGeneration(id string) Generation {
	var found bool = false
	var gen Generation
	for _, g := range generations {

		if id == fmt.Sprintf("%d", g.Id) {
			found = true
			gen = (*g)
		}
	}
	if !found {
		panic("No Generation matches.")
	}
	return gen
}

func GetGenerationMaxId() int {
	var maxId int = -1
	for _, g := range generations {
		if g.Id > maxId {
			maxId = g.Id
		}
	}
	return maxId
}

func GetGenerationsByCompany(id string) []Generation {
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

func GetVTuberMaxId() int {
	var maxId int = -1
	for _, vt := range vtubers {
		if vt.Id > maxId {
			maxId = vt.Id
		}
	}
	return maxId
}

func GetVTubersByCompany(id string) []VTuber {
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

func MakeCompanyOptions() string {
	var options string = ""
	for _, company := range companies {
		options = fmt.Sprintf("%s<option value=\"%d\">%s</option>", options, company.Id, company.Name)
	}
	return options
}

func MakeGenerationOptions() string {
	var options string = ""
	for _, generation := range generations {
		options = fmt.Sprintf("%s<option value=\"%d\">%s</option>", options, generation.Id, generation.Name)
	}
	return options
}

func MakeHtml(msg string) string {
	return fmt.Sprintf("<html><head><title>Server</title></head><body>%s<hr><p><a href=\"/\">Home</a></body></html>", msg)
}

func DisplayPage(w http.ResponseWriter, c string) {
	w.Header().Set("Content-Type", "text/html")
	io.WriteString(w, c)
}
