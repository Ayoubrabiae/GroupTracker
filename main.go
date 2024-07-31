package main

import (
	"fmt"
	"log"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
	"groupie-tracker/handlers"
)

func main() {
	var err error

	err = funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = funcs.GetAndParse(data.MainData.Artists, &data.Artist)
	if err != nil {
		fmt.Println(err)
		return
	}

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists/", handlers.ProfileHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("http://localhost:8050/")
	fmt.Println(-2 + 5)
	log.Fatal(http.ListenAndServe(":8050", http.DefaultServeMux))
}
