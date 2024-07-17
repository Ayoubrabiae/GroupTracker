package main

import (
	"fmt"

	"groupie-tracker/data"
	"groupie-tracker/funcs"
)

func main() {
	funcs.GetAndParse("https://groupietrackers.herokuapp.com/api", &data.MainData)

	fmt.Println(data.MainData)
}