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
		http.Error(w, "Bad Reaquest", http.StatusBadRequest)
		return
	}

	matches := IdPath.FindStringSubmatch(r.URL.Path)

	if len(matches) < 1 {
		http.Error(w, "Page Not Found", http.StatusNotFound)
		return
	}

	userId, err := strconv.Atoi(matches[1])
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	if userId >= len(data.Artist) {
		http.Error(w, "Page Not Found", http.StatusBadRequest)
		return
	}

	tmp, err := template.ParseFiles("./static/pages/profile.html")
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

	userId--

	profileData := struct {
		Artist    data.ArtistType
		Locations []string
		Dates     []string
		Relations map[string][]string
	}{
		Artist:    data.Artist[userId],
		Locations: data.Locations.Index[userId].Locations,
		Dates:     data.Dates.Index[userId].Dates,
		Relations: data.Relations.Index[userId].DatesLocations,
	}

	fmt.Println(data.Relations.Index[userId])

	err = tmp.Execute(w, profileData)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		fmt.Println("When we excute")
		return
	}
}
