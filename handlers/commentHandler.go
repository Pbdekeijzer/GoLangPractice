package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"strconv"

	"github.com/gorilla/mux"
	"github.com/pbdekeijzer/GoLangPractice/datastorage"
	"github.com/pbdekeijzer/GoLangPractice/models"
)

func GetAllComments(w http.ResponseWriter, r *http.Request) {
	comments, err := json.MarshalIndent(datastorage.GetAllComments(), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(comments)
}

func GetIssueComments(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	comments, err := json.MarshalIndent(datastorage.GetCommentsByIssue(id), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(comments)
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var comment models.Comment

	id, err := strconv.Atoi(vars["id"])

	rBody, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(rBody, &comment)

	if err != nil || error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	datastorage.CreateComment(comment, id)
}
