package handlers

import (
	"fmt"
	"net/http"
)

// IndexHandler is the root at "/"
func IndexHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the JSON restAPI!")
}
