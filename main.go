package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"groupie/data"
	"groupie/funcs"
	"groupie/handlers"
)

func main() {
	res, err := funcs.GetJsonData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	json.Unmarshal([]byte(res), &data.Groups)

	mux := http.NewServeMux()
	stylizeFolder := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static/", stylizeFolder))
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/groups/", handlers.ProfileHandler)

	server := http.Server{
		Addr:    ":4561",
		Handler: mux,
	}

	fmt.Println("http://localhost:4561")

	server.ListenAndServe()
}
