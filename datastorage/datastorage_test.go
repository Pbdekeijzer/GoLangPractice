package datastorage

import (
	"testing"

	"github.com/pbdekeijzer/GoLangPractice/models"
)

// setup issue object
func setUpIssue() models.Issue {
	issue := models.Issue{IssueContent: "Test issue creation", Status: "Busy", Comments: nil}
	return CreateIssue(issue)
}

// setup comment object
func setUpComment() models.Comment {
	comment := models.Comment{Content: "This is the content", OnIssue: "Test issue creation"}
	return CreateComment(comment, 1)
}

func TestCreateIssue(t *testing.T) {
	issue := setUpIssue()

	newIssue := GetIssue(issue.ID)

	if issue.ID != newIssue.ID {
		t.Error("The issue returned is not the same as the initial issue")
	}
}

func TestDeleteIssue(t *testing.T) {
	issue := setUpIssue()

	DeleteIssue(issue.ID)

	if GetIssue(issue.ID).IssueContent == "Test issue creation" {
		t.Error("Not succesfully deleted the issue")
	}
}

func TestEditIssue(t *testing.T) {
	issue := setUpIssue()

	updatedIssue := models.Issue{ID: issue.ID, IssueContent: "Test issue creation", Status: "Done", Comments: nil}
	EditIssue(updatedIssue)

	newIssue := GetIssue(issue.ID)

	if newIssue.Status != "Done" {
		t.Error("Not succesfully updated the issue")
	}
}

func TestGetAllIssues(t *testing.T) {
	setUpIssue()

	allIssues := GetAllIssues()

	if len(allIssues) == 0 {
		t.Error("The function returns an empty slice when trying to get all issues")
	}
}

func TestGetIssue(t *testing.T) {
	issue := setUpIssue()

	newissue := GetIssue(issue.ID)

	if newissue.ID != issue.ID {
		t.Error("The returned ID is not the same as the created one")
	}
}

func TestCreateComment(t *testing.T) {
	comment := setUpComment()
	fail := true
	comments := GetAllComments()

	for _, com := range comments {
		if comment.CommentID == com.CommentID {
			fail = false
		}
	}

	if fail == true {
		t.Error("The comment was not created")
	}
}

func TestDeleteComment(t *testing.T) {
	comment := setUpComment()
	fail := false

	issue := GetIssue(1)
	DeleteIssue(issue.ID)

	comments := GetAllComments()

	for _, com := range comments {
		if com.CommentID == comment.CommentID {
			fail = true
		}
	}

	if fail == true {
		t.Error("The comment was not deleted")
	}
}

func TestGetAllComments(t *testing.T) {
	setUpComment()

	comments := GetAllComments()

	if len(comments) == 0 {
		t.Error("The function returns an empty slice when trying to get all issues")
	}
}

func TestGetCommentsByIssue(t *testing.T) {
	comments := GetCommentsByIssue(1)
	issue := GetIssue(1)
	if len(comments) != len(issue.Comments) {
		t.Error("The functions doesn't return the correct comments")
	}
}
