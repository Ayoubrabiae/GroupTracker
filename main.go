package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Artist struct {
	Id           int    `json:"id"`
	Image        string `json:"image"`
	Name         string `json:"name"`
	Members      string `json:"members"`
	CreationDate int    `json:"creationDate"`
	FirstAlbum   string `json:"firstAlbum"`
	Locations    string `json:"locations"`
	ConcertDates string `json:"concerDates"`
	Relations    string `json:"relations"`
}

type Locations struct {
	
}

func main() {
	url := "https://groupietrackers.herokuapp.com/api/artists"
	var artist []Artist
	D, _ := http.Get(url)
	json.NewDecoder(D.Body).Decode(&artist)
	fmt.Println(artist[5].Name)
}
