package main

import (
	"encoding/json"
	"fmt"
	"groupie/data"
	"groupie/funcs"
	"groupie/handlers"
	"net/http"
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

	server := http.Server{
		Addr:    ":4561",
		Handler: mux,
	}

	server.ListenAndServe()
}
