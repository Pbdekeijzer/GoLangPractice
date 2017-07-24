package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"encoding/json"

	"bytes"

	"github.com/pbdekeijzer/GoLangPractice/models"
	"github.com/stretchr/testify/assert"
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
				"cid": 1,
				"content": "This api needs some work",
				"issue": 3
			},
			{
				"cid": 2,
				"content": "This api is not public",
				"issue": 3
			}
		]
	},
	{
		"id": 4,
		"issuecontent": "Add Docker",
		"status": "ToDo",
		"comments": [
			{
				"cid": 3,
				"content": "No authentication!? UNSAFE",
				"issue": 4
			}
		]
	}
]`

	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}

func TestPostIssueHandler(t *testing.T) {
	issue := models.Issue{
		ID:           1,
		IssueContent: "Test post issue",
		Status:       "Pending",
		Comments:     nil,
	}

	jsonIssue, _ := json.Marshal(issue)
	req := httptest.NewRequest("POST", "/issue", bytes.NewBuffer(jsonIssue))

	response := httptest.NewRecorder()

	handler := http.HandlerFunc(PostIssueHandler)
	handler.ServeHTTP(response, req)
	assert.Equal(t, 201, response.Code, "Response 201 is expected")
}
