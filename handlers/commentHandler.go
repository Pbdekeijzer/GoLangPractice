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

// GetAllCommentsHandler returns all comments in json format
func GetAllCommentsHandler(w http.ResponseWriter, r *http.Request) {
	comments, err := json.MarshalIndent(datastorage.GetAllComments(), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(comments)
}

// GetIssueCommentsHandler returns the comments of an issue in json format
// If the issue doesn't exist, comments = [], with len 2
// If len < 3, return 404, else return the comments or null if the issue has no comments
func GetIssueCommentsHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	comments, err := json.MarshalIndent(datastorage.GetCommentsByIssue(id), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if len(comments) < 3 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.Header().Set("Content-Type", "application/json")
		w.Write(comments)
	}
}

// PostCommentHandler unmarshals json body to a models.Comment and adds the comment to the datastorage
func PostCommentHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var comment models.Comment

	id, err := strconv.Atoi(vars["id"])

	rBody, _ := ioutil.ReadAll(r.Body)
	error := json.Unmarshal(rBody, &comment)

	if err != nil || error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	r.Body.Close()

	datastorage.CreateComment(comment, id)

	w.WriteHeader(http.StatusCreated)
}
