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

func PostIssue(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue
	rBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//add an issue to the data storage
	datastorage.CreateIssue(issue)

	//just for test
	newissue, err := json.Marshal(issue)
	w.Header().Set("Content-Type", "application/json")
	w.Write(newissue)
}

func GetAllIssues(w http.ResponseWriter, r *http.Request) {
	issues, err := json.MarshalIndent(datastorage.GetAllIssues(), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(issues)
}

func GetIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	issue, err := json.MarshalIndent(datastorage.GetIssue(id), "", "\t")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(issue)
}

func DeleteIssue(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	datastorage.DeleteIssue(id)
	w.WriteHeader(http.StatusNoContent)
}

func PatchIssue(w http.ResponseWriter, r *http.Request) {
	var issue models.Issue

	rBody, _ := ioutil.ReadAll(r.Body)
	err := json.Unmarshal(rBody, &issue)

	datastorage.EditIssue(issue)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	//apparently doesn't automatically return 200 on patch
	w.WriteHeader(http.StatusOK)
}
