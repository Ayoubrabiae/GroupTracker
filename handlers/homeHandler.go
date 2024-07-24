package handlers

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./static/pages/index.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusNotFound)
		fmt.Println("When we parse the index.html")
		return
	}

	tmp, err = tmp.ParseGlob("./static/templates/*.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusNotFound)
		fmt.Println("When we parse all templates")
		return
	}

	err = funcs.GetAndParse(data.MainData.Artists, &data.Artist)
	if err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
	}

	err = tmp.Execute(w, data.Artist)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusNotFound)
		fmt.Println("When we excute the html")
		return
	}
}
