package handlers

import (
	"fmt"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the JSON restAPI!")
}

func MoreIndex(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the JSON restAPI!")
}
