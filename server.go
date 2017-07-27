package main

import (
	"log"
	"net/http"

	"github.com/pbdekeijzer/GoLangPractice/router"
	"github.com/urfave/negroni"
)

func main() {

	router := router.NewRouter()

	n := negroni.Classic()
	n.UseHandler(router)

	log.Fatal(http.ListenAndServe(":8080", nil))
}
