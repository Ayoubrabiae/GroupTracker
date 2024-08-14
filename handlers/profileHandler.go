package handlers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"

	"groupie-tracker/data"
)

var IdPath = regexp.MustCompile(`^/artists/(\d+)$`)

func ProfileHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		ErrorHandler(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	matches := IdPath.FindStringSubmatch(r.URL.Path)

	if len(matches) < 2 {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	id := matches[1]

	if id == "360" {
		tmp, err := template.ParseFiles("./pages/profile.html")
		if err != nil {
			ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
			fmt.Println("When we parse the index.html")
			return
		}

		profileData, err := data.ArtistInfo{
			Artist: data.ArtistType{
				Id:           360,
				Image:        "/static/images/cheb.jpg",
				Name:         "Cheb L3arbi",
				Members:      []string{"Cheb L3arebi"},
				CreationDate: 2020,
				FirstAlbum:   "2019-2-30",
			},
			Locations: data.LocationsType{
				Id:        360,
				Locations: []string{"Darhom", "Ras-Derb", "New-York", "LV", "Asfi"},
				Dates:     "kljdfhqsdhlfkhgqkjsdgf",
			},
			Dates: data.DatesType{
				Id: 360,
				Dates: []string{
					"2019-05-20",
					"2019-05-21",
					"2019-06-20",
					"2019-02-31",
					"1942-09-14",
					"2019-05-20",
					"2006-05-12",
				},
			},
			Relations: data.RelationsType{
				Id: 360,
				DatesLocations: map[string][]string{
					"Darhom": {
						"2019-05-20",
						"2019-05-21",
						"2019-06-20",
						"2019-02-31",
						"1942-09-14",
					},
					"Ras-Derb": {
						"1942-09-14",
						"2019-05-20",
						"2006-05-12",
					},
				},
			},
		}, nil

		err = tmp.Execute(w, profileData)
		if err != nil {
			fmt.Println("When we excute", err)
			return
		}
		return
	}

	ok, err := data.IdCheck(id)
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		return
	}
	if !ok {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	tmp, err := template.ParseFiles("./pages/profile.html")
	if err != nil {
		ErrorHandler(w, "Internal Server error", http.StatusInternalServerError)
		fmt.Println("When we parse the index.html")
		return
	}

	profileData, err := data.LoadArtistData(id)
	if err != nil {
		ErrorHandler(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	err = tmp.Execute(w, profileData)
	if err != nil {
		fmt.Println("When we excute", err)
		return
	}
}
