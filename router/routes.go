package router

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
	Route{"Index", "GET", "/", handlers.IndexHandler},

	//Issue routes
	Route{"GetAllIssues", "GET", "/issues", handlers.GetAllIssuesHandler},
	Route{"GetIssue", "GET", "/issue/{id}", handlers.GetIssueHandler},
	Route{"PostIssue", "POST", "/issue", handlers.PostIssueHandler},
	Route{"PatchIssue", "PATCH", "/issue/{id}", handlers.PatchIssueHandler},
	Route{"DeleteIssue", "DELETE", "/issue/{id}", handlers.DeleteIssueHandler},

	// //Comment routes
	Route{"PostComment", "POST", "/issue/{id}/comment", handlers.PostCommentHandler},
	Route{"GetAllComments", "GET", "/comments", handlers.GetAllCommentsHandler},
	Route{"GetIssueComments", "GET", "/issue/{id}/comments", handlers.GetIssueCommentsHandler},
}
