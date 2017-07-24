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
		"issue": 3
	},
	{
		"cid": 2,
		"content": "This api is not public",
		"issue": 3
	},
	{
		"cid": 3,
		"content": "No authentication!? UNSAFE",
		"issue": 4
	}
]`

	if responseRecorder.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v, want %v", responseRecorder.Body.String(), expected)
	}
}

// UNIT TEST NOT WORKING WITH URL PARAMETERS, MAYBE FIX LATER
// func TestCommentIssueHandler(t *testing.T) {
// 	comment := models.Comment{
// 		CommentID: 1,
// 		Content:   "Comment content",
// 		OnIssue:   "Add Authentication",
// 	}

// 	jsonComment, _ := json.Marshal(comment)

// 	req := httptest.NewRequest("POST", "/issue/1/comment", bytes.NewBuffer(jsonComment))

// 	response := httptest.NewRecorder()

// 	handler := http.HandlerFunc(PostCommentHandler)
// 	handler.ServeHTTP(response, req)
// 	assert.Equal(t, 201, response.Code, "Response 201 is expected")
// }
