package main

import (
	"fmt"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artist []Artist
	D, _ := http.Get(url)
	json.NewDecoder(D.Body).Decode(&artist)
	fmt.Println(artist[5].Name)
}
