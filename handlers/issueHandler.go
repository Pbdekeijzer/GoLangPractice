package handlers

import (
	"net/http"
	"strconv"

	"encoding/json"

	"io/ioutil"

	"github.com/gorilla/mux"
	"github.com/pbdekeijzer/GoLangPractice/datastorage"
	"github.com/pbdekeijzer/GoLangPractice/models"
)

// PostIssueHandler unmarshals json body to models.Issue and adds the issue to the datastorage
func PostIssueHandler(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue
	rBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datastorage.CreateIssue(issue)
	w.WriteHeader(http.StatusCreated)
}

// GetAllIssuesHandler returns json containing all issues
func GetAllIssuesHandler(w http.ResponseWriter, r *http.Request) {
	issues, err := json.MarshalIndent(datastorage.GetAllIssues(), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(issues)
}

// GetIssueHandler returns json containing a single issue
func GetIssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	issue := datastorage.GetIssue(id)
	if issue.ID == 0 {
		w.WriteHeader(http.StatusNotFound)
	} else {
		newissue, error := json.MarshalIndent(issue, "", "\t")
		if error != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(newissue)
	}

}

// DeleteIssueHandler deletes an issue, based on the is in the uri
func DeleteIssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	err = datastorage.DeleteIssue(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}

}

// PutIssueHandler asks for an existing issue
// It checks the ID and replaces the object in the issues slice with the new object
// return 204 if function was succesful
// TO DO: - Implements this more elegantly without the need for an issue
func PutIssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	var issue models.Issue

	rBody, _ := ioutil.ReadAll(r.Body)
	err = json.Unmarshal(rBody, &issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	r.Body.Close()

	err = datastorage.EditIssue(issue, id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
	} else {
		w.WriteHeader(http.StatusNoContent)
	}
}
