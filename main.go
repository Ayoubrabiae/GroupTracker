package main

import (
	"fmt"
	"log"
	"net/http"
	//"os"

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

	//port := os.Getenv("PORT")

	//if port == "" {
		port := "8080"
	//}

	http.HandleFunc("/", handlers.HomeHandler)
	http.HandleFunc("/artists/", handlers.ProfileHandler)
	http.HandleFunc("/about", handlers.AboutHandler)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	fmt.Println("http://localhost:" + port)
	log.Fatal(http.ListenAndServe(":"+port, http.DefaultServeMux))
}
