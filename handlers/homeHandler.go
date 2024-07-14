package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie/data"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "404 - Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != "GET" {
		http.Error(w, "400 - Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("static/index.html")
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		fmt.Println("Error when we parse index", err)
		return
	}

	tmpl, err = tmpl.ParseGlob("static/templates/*.html")
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		fmt.Println("Error when we parse templates", err)
		return
	}

	tmpl.Execute(w, data.Groups)
}
