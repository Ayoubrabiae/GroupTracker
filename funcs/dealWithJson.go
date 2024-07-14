package funcs

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"groupie/data"
)

func GetJsonData(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	resData, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resData), nil
}

func ParseJsonData(v interface{}, url string) {
	res, err := GetJsonData(url)
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	err = json.Unmarshal([]byte(res), v)
	if err != nil {
		fmt.Println("Error in parsing:", err)
	}
}

func FillData() {
	ParseJsonData(&data.Groups, "https://groupietrackers.herokuapp.com/api/artists")
	ParseJsonData(&data.LocationsHolder, "https://groupietrackers.herokuapp.com/api/locations")
	ParseJsonData(&data.DatesHolder, "https://groupietrackers.herokuapp.com/api/dates")
	ParseJsonData(&data.RelationsHolder, "https://groupietrackers.herokuapp.com/api/relation")
}
