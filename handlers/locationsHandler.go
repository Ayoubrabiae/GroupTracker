package handlers

import (
	"fmt"
	"net/http"
	"regexp"
)

func LocationsHandler(w http.ResponseWriter, r *http.Request) {
	re := regexp.MustCompile(`/locations/[w|-]+`)

	matches := re.FindStringSubmatch(r.URL.Path)
	if len(matches) < 2 {
		ErrorHandler(w, "Page Not Found", http.StatusNotFound)
		return
	}

	place := matches[1]

	fmt.Println(place)

	http.Redirect(w, r, "/", http.StatusSeeOther)
}
