package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllIssuesHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/issues", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllIssuesHandler)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[
	{
		"id": 1,
		"issuecontent": "Test create and edit functions",
		"status": "Pending",
		"comments": null
	},
	{
		"id": 2,
		"issuecontent": "Retreive data from datastorage",
		"status": "Pending",
		"comments": null
	},
	{
		"id": 3,
		"issuecontent": "Add Authentication",
		"status": "ToDo",
		"comments": [
			{
				"cid": 2,
				"content": "This api is not public",
				"OnIssue": ""
			},
			{
				"cid": 3,
				"content": "No authentication!? UNSAFE",
				"OnIssue": ""
			}
		]
	},
	{
		"id": 4,
		"issuecontent": "Add Docker",
		"status": "ToDo",
		"comments": [
			{
				"cid": 1,
				"content": "This api needs some work",
				"OnIssue": ""
			}
		]
	}
]`
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}

// func TestGetIssueHandler(t *testing.T) {
// 	req, err := http.NewRequest("GET", "/issue/1", nil)
// 	if err != nil {
// 		t.Fatal(err)
// 	}

// 	responseRecorder := httptest.NewRecorder()
// 	handler := http.HandlerFunc(GetIssueHandler)

// 	handler.ServeHTTP(responseRecorder, req)

// 	if status := responseRecorder.Code; status != http.StatusOK {
// 		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
// 	}

// 	expected := `{
// 	"id": 1,
// 	"issuecontent": "Test create and edit functions",
// 	"status": "Pending",
// 	"comments": null
// }`
// 	if responseRecorder.Body.String() != expected {
// 		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
// 	}
// }
