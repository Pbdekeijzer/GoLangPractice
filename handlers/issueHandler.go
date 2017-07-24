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

// PostIssue unmarshals json body to models.Issue and adds the issue to the datastorage
func PostIssueHandler(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue
	rBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	datastorage.CreateIssue(issue)
}

// GetAllIssues returns json containing all issues
func GetAllIssuesHandler(w http.ResponseWriter, r *http.Request) {
	issues, err := json.MarshalIndent(datastorage.GetAllIssues(), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(issues)
}

// GetIssue returns json containing a single issue
func GetIssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	issue, err := json.MarshalIndent(datastorage.GetIssue(id), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(issue)
}

// DeleteIssue deletes an issue, based on the is in the uri
func DeleteIssueHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	datastorage.DeleteIssue(id)
	w.WriteHeader(http.StatusNoContent)
}

// PatchIssue asks for an existing issue
// It checks the ID and modifies the rest of the data
// return 200 if function was succesful (patch apparently doesn't do this automatically)
// TO DO: Implements this more elegantly without the need for an issue
func PatchIssueHandler(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue

	rBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &issue)

	datastorage.EditIssue(issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.WriteHeader(http.StatusOK)
}
