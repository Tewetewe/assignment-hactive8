package main

import (
	"fmt"
	"html/template"
	"net/http"
)

type Biodata struct {
	Name  string
	Email string
}

var PORT = ":8080"

var biodata = []Biodata{}

func main() {
	http.HandleFunc("/post", postBiodata)
	http.HandleFunc("/", index)

	fmt.Println("Application is listening on port", PORT)

	http.ListenAndServe(PORT, nil)

}

func index(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		var tmpl = template.Must(template.New("loginForm").ParseFiles("template.html"))

		var err = tmpl.Execute(w, nil)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "Invalid Method", http.StatusBadRequest)
}

func postBiodata(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var tmpl = template.Must(template.New("home").ParseFiles("template.html"))

		if err := r.ParseForm(); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		name := r.FormValue("name")
		email := r.FormValue("email")

		newBiodata := Biodata{
			Name:  name,
			Email: email,
		}

		biodata = append(biodata, newBiodata)

		if err := tmpl.Execute(w, biodata); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	http.Error(w, "", http.StatusBadRequest)

}
