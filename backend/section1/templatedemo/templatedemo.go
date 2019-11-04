package main

import (
	"html/template"
	"log"
	"net/http"
)

// Gopher type
type Gopher struct {
	Name string
}

func main() {
	http.HandleFunc("/hello-gopher", helloGopherHandler)
	http.ListenAndServe(":8080", nil)
}

// Handler for hello-gopher route
func helloGopherHandler(w http.ResponseWriter, r *http.Request) {
	gopherName := "Gopher"
	queryParam := r.URL.Query().Get("gophername")
	if queryParam != "" {
		gopherName = queryParam
	}

	gopher := Gopher{Name: gopherName}
	renderTemplate(w, "./templates/greeting.html", gopher)
}

// template rendering function
func renderTemplate(w http.ResponseWriter, templateFilePath string, data interface{}) {
	t, err := template.ParseFiles(templateFilePath)
	if err != nil {
		log.Fatal("Error while parsing the template ", err)
	}
	t.Execute(w, data)
}
