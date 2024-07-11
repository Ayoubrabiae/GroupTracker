package main

import (
	"encoding/json"
	"fmt"
	"groupie/data"
	"groupie/funcs"
)

func main() {
	res, err := funcs.GetJsonData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	var Groups []data.Group

	json.Unmarshal([]byte(res), &Groups)

	for _, gr := range Groups {
		fmt.Println(gr)
		fmt.Println()
	}
}
