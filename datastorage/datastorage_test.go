package datastorage

import (
	"testing"

	"fmt"

	"github.com/pbdekeijzer/GoLangPractice/models"
	"github.com/stretchr/testify/assert"
)

var id = 1

// setup issue object
func setUpIssue() models.Issue {
	issue := models.Issue{IssueContent: "Test issue creation", Status: "Busy", Comments: nil}
	return CreateIssue(issue)
}

// setup comment object
func setUpComment() (models.Comment, error) {
	comment := models.Comment{Content: "This is the content", OnIssue: id}
	return CreateComment(comment, id)
}

func TestCreateIssue(t *testing.T) {
	issue := setUpIssue()

	newIssue := GetIssue(issue.ID)

	assert.Equal(t, issue.ID, newIssue.ID, fmt.Sprintf("The ID should be %v", issue.ID))

}

func TestDeleteIssue(t *testing.T) {
	issue := setUpIssue()

	DeleteIssue(issue.ID)

	assert.NotEqual(t, "Test issue creation", GetIssue(issue.ID).IssueContent, fmt.Sprintf("Issue content should not be: Test issue creation"))
}

func TestEditIssue(t *testing.T) {
	issue := setUpIssue()

	updatedIssue := models.Issue{ID: issue.ID, IssueContent: "Test issue creation", Status: "Done", Comments: nil}
	EditIssue(updatedIssue, issue.ID)

	newIssue := GetIssue(issue.ID)

	assert.Equal(t, "Done", newIssue.Status, fmt.Sprintf("The issue status should be: Done"))
}

func TestGetAllIssues(t *testing.T) {
	setUpIssue()

	allIssues := GetAllIssues()

	assert.NotEqual(t, 0, len(allIssues), "Should not return a slice of length 0")
}

func TestGetIssue(t *testing.T) {
	issue := setUpIssue()

	newissue := GetIssue(issue.ID)

	assert.Equal(t, issue.ID, newissue.ID, fmt.Sprintf("The issue ID should be %v", issue.ID))
}

func TestCreateComment(t *testing.T) {
	comment, err := setUpComment()
	fail := true

	if err != nil {

	}
	comments := GetAllComments()

	for _, com := range comments {
		if comment.CommentID == com.CommentID {
			fail = false
		}
	}

	//if not false, newly created comment not found in list of comments
	assert.Equal(t, false, fail, "fail should be false")
}

func TestDeleteComment(t *testing.T) {
	comment, err := setUpComment()
	fail := false
	if err != nil {
		issue := GetIssue(1)
		DeleteIssue(issue.ID)

		comments := GetAllComments()

		for _, com := range comments {
			if com.CommentID == comment.CommentID {
				fail = true
			}
		}
	}

	assert.Equal(t, false, fail, "Fail is false when comment is succesfully deleted")
}

func TestGetAllComments(t *testing.T) {
	setUpComment()

	comments := GetAllComments()

	assert.NotEqual(t, 0, len(comments), "Should not return a slice of length 0")
}

func TestGetCommentsByIssue(t *testing.T) {
	comments := GetCommentsByIssue(id)
	issue := GetIssue(id)

	assert.Equal(t, len(issue.Comments), len(comments), "The length of the slice should be the same")
}
