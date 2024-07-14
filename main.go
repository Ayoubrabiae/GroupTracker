package main

import (
	"fmt"
	"net/http"

	"groupie/funcs"
	"groupie/handlers"
)

func main() {
	funcs.FillData()

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
