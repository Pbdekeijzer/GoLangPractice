package main

import (
	"net/http"

	"github.com/pbdekeijzer/restAPI/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	//IndexSite
	Route{"Index", "GET", "/", handlers.Index},

	//Issue routes
	// Route{"GetAllIssues", "GET", "/issues", GetAllIssues},
	// Route{"GetIssue", "GET", "issue/{id}", GetIssue},
	// Route{"CreateIssue", "POST", "issue", CreateIssue},
	// Route{"UpdateIssue", "PATCH", "issue/{id}", UpdateIssue},
	// Route{"DeleteIssue", "DELETE", "issue/{id}", DeleteIssue},

	// //Comment routes
	// Route{"CreateComment", "POST", "issue/{id}/comment", CreateComment},
	// Route{"GetAllComments", "GET", "/comments", GetAllComments},
	// Route{"GetIssueComments", "GET", "issue/{id}/comments", GetIssueComments},
}
