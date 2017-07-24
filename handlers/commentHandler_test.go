package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestGetAllCommentsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetAllCommentsHandler)

	handler.ServeHTTP(responseRecorder, req)

	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[
	{
		"cid": 1,
		"content": "This api needs some work",
		"OnIssue": "Add Docker"
	},
	{
		"cid": 2,
		"content": "This api is not public",
		"OnIssue": "Add Authentication"
	},
	{
		"cid": 3,
		"content": "No authentication!? UNSAFE",
		"OnIssue": "Add Authentication"
	}
]`
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}

func TestGetIssueCommentsHandler(t *testing.T) {
	req, err := http.NewRequest("GET", "/issue/4/comments", nil)
	if err != nil {
		t.Fatal(err)
	}

	responseRecorder := httptest.NewRecorder()
	handler := http.HandlerFunc(GetIssueCommentsHandler)

	handler.ServeHTTP(responseRecorder, req)
	if status := responseRecorder.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v, want %v", status, http.StatusOK)
	}

	expected := `[
	{
		"cid": 1,
		"content": "This api needs some work",
		"OnIssue": ""
	}
]`
	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}
