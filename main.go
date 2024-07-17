package main

import (
	"fmt"
	"html/template"
	"net/http"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

// var tp1 *template.Template
func main() {
	// template.ParseGlob("./Template/*.html")
	http.HandleFunc("/", main_Page)
	http.ListenAndServe(":8081", nil)
	// fmt.Println(data.Artist)
}

func main_Page(w http.ResponseWriter, r *http.Request) {
	err := funcs.GetAndParse("https://groupietrackers.herokuapp.com/api/artists", &data.Artist)
	if err != nil {
		fmt.Println(err)
	}
	tp1, _ := template.ParseFiles("Template/Cards.html")
	err = tp1.ExecuteTemplate(w, "Cards.html", data.Artist)
	fmt.Println(err)
}
