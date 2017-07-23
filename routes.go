package main

import (
	"net/http"

	"github.com/pbdekeijzer/GoLangPractice/handlers"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{

	//Index
	Route{"Index", "GET", "/", handlers.Index},

	//Issue routes
	Route{"GetAllIssues", "GET", "/issues", handlers.GetAllIssues},
	Route{"GetIssue", "GET", "/issue/{id}", handlers.GetIssue},
	Route{"PostIssue", "POST", "/issue", handlers.PostIssue},
	Route{"PatchIssue", "PATCH", "/issue/{id}", handlers.PatchIssue},
	Route{"DeleteIssue", "DELETE", "/issue/{id}", handlers.DeleteIssue},

	// //Comment routes
	Route{"PostComment", "POST", "/issue/{id}/comment", handlers.PostComment},
	Route{"GetAllComments", "GET", "/comments", handlers.GetAllComments},
	Route{"GetIssueComments", "GET", "/issue/{id}/comments", handlers.GetIssueComments},
}
