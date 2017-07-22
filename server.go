package main

import (
	"log"
	"net/http"
)

// func Index(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintf(w, "Helloossso, %q", html.EscapeString(r.URL.Path))
// }

func main() {

	router := NewRouter()
	log.Fatal(http.ListenAndServe(":8080", router))
}
