package funcs

import (
	"encoding/json"
	"net/http"
)

func GetAndParse(url string, v interface{}) error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}

	err = json.NewDecoder(res.Body).Decode(v)
	if err != nil {
		return nil
	}

	return nil
}
