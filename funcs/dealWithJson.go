package funcs

import (
	"encoding/json"
	"net/http"
)

func GetAndParse(s string, any interface{}) error {
	D, err := http.Get(s)
	if err != nil {
		return err
	}

	err = json.NewDecoder(D.Body).Decode(&any)
	if err != nil {
		return err
	}
	return nil
}
