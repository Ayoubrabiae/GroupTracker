package handlers

import (
	"fmt"
	"groupie/data"
	"html/template"
	"net/http"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
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
