package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"groupie-tracker/data"
)

var IdPath = regexp.MustCompile(`^/artists/(\d+)$`)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	matches := IdPath.FindStringSubmatch(r.URL.Path)

	if len(matches) < 1 {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	id := matches[1]

	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if idNum > len(data.Artist) {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("./static/pages/profile.html")
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

	profileData, err := data.LoadArtistData(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	err = tmp.Execute(w, profileData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("When we excute", err)
		return
	}
}
