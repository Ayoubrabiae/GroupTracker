package funcs

import (
	"io"
	"net/http"
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
