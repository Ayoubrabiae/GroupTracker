package main

import (
	"log"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func main() {
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api/artists", &data.Artist)
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api/locations", &data.Locations)
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api/dates", &data.Dates)
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api/relation", &data.Relations)

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists/", handlers.ProfileHandler)
	log.Fatal(http.ListenAndServe(":8080", http.DefaultServeMux))
}
