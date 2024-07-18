package handlers

import (
	"fmt"
	"net/http"
	"text/template"

	"groupie-tracker/data"
)

var tp1 *template.Template

func Main_Page(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not found", 404)
		return
	}
	tp1, err := template.ParseFiles("static/templates/card.html")
	if err != nil {
		fmt.Println(err)
		return
	}

	err2 := tp1.ExecuteTemplate(w, "card.html", data.Artist)
	if err2 != nil {
		fmt.Println(err2)
		return
	}
}
