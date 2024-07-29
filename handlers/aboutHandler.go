package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

func AboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	tmp, err := template.ParseFiles("./static/pages/about.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse the index.html")
		return
	}

	tmp, err = tmp.ParseGlob("./static/templates/*.html")
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse all templates")
		return
	}

	err = tmp.Execute(w, nil)
	if err != nil {
		http.Error(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we excute")
		return
	}
}
