package main

import (
	"fmt"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func main() {
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)
	funcs.GetAndParse(data.MainData.Artists, &data.Artist)
	funcs.GetAndParse(data.MainData.Relations, &data.Relations)
	funcs.GetAndParse(data.MainData.Dates, &data.Dates)
	funcs.GetAndParse(data.MainData.Locations, &data.Locations)

	mux := http.NewServeMux()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/artists/", handlers.ProfileHandler)
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	server := http.Server{
		Addr:    ":8050",
		Handler: mux,
	}

	fmt.Println("http://localhost:8050")
	server.ListenAndServe()
}
