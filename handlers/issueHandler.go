package handlers

import (
	"fmt"
	"net/http"

	"github.com/pbdekeijzer/GoLangPractice/datastorage"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to the JSON restAPI!")
	datastorage.DeleteIssue(2)
}

// func GetIssue(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// }
