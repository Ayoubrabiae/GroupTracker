package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"groupie/data"
)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	pathParts := strings.Split(r.URL.Path, "/")

	if len(pathParts) != 3 {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	id := pathParts[2]

	idNum, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	if idNum > len(data.Groups) && idNum < len(data.Groups) {
		http.Error(w, "500 - Internal Server Error", http.StatusInternalServerError)
		return
	}

	profileData := data.ProfileData{
		Group:     data.Groups[idNum-1],
		Locations: data.LocationsHolder.Index[idNum-1].Locations,
		Dates:     data.DatesHolder.Index[idNum-1].Dates,
		Relations: data.RelationsHolder.Index[idNum-1].DatesLocations,
	}

	tmpl, err := template.ParseGlob("static/templates/*.html")
	if err != nil {
		http.Error(w, "Server Error 500", http.StatusInternalServerError)
		fmt.Println("Error when we parse templates", err)
		return
	}

	tmpl.ExecuteTemplate(w, "profile", profileData)
}
