package handlers

import (
	"fmt"
	"net/http"
	"strconv"
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

func Infos(w http.ResponseWriter, r *http.Request) {
	var digit string
	tp1,err := template.ParseGlob("static/pages/profile.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	digit = r.URL.Path[9:]
	d, err1 := strconv.Atoi(digit)
if err1 != nil {
	fmt.Println(err1)
	return
}
	d--
	p := data.AllinfosType {
		ArtistT:data.Artist[d],
		LocationsT:data.Locations[d],
		DatesT:data.Dates[d],
		RelationsT:data.Relations[d],
	 }

	err2 := tp1.ExecuteTemplate(w, "profile.html", p)
	if err2 != nil {
		 fmt.Println(err2)
		return
	}
}
