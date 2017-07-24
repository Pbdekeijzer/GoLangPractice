package datastorage

import (
	"fmt"

	"github.com/pbdekeijzer/GoLangPractice/models"
)

// issues is where the data for all the issues is stored
// sid is used for auto incrementing the id
var issues models.Issues
var sid int

// CreateIssue creates an issue and appends the issue to the issues slice
// Auto increments the ID, so any id gives is transformed into the incremented id given by this function
func CreateIssue(issue models.Issue) models.Issue {
	sid++
	issue.ID = sid
	issues = append(issues, issue)
	return issue
}

// DeleteIssue deletes an issue from the issues slice
// Also deletes all the comments the issue had
func DeleteIssue(id int) error {

	for i, s := range issues {
		if s.ID == id {

			if len(issues[i].Comments) != 0 {
				for _, comment := range issues[i].Comments {
					DeleteComment(comment.CommentID)
				}
			}

			issues = append(issues[:i], issues[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("There is no issue with id %d to delete", id)
}

// EditIssue simply modifies an existing issue, based on the issue.ID
func EditIssue(issue models.Issue, sid int) error {
	for i := 0; i < len(issues); i++ {
		if issues[i].ID == sid {
			issues[i] = issue
			return nil
		}
	}
	return fmt.Errorf("Not able to edit issue with id: %d", sid)
}

// GetAllIssues returns a slice of the struct models.Issue
func GetAllIssues() models.Issues {
	return issues
}

// GetIssue returns a single struct of models.Issue
func GetIssue(id int) models.Issue {
	for i := 0; i < len(issues); i++ {
		if issues[i].ID == id {
			return issues[i]
		}
	}
	// else return empty Issue
	return models.Issue{}
}

//Comment
var comments models.Comments
var cid int

// CreateComment and appends the comment to the comments of the corresponding issue
// Auto increment the comment id and transform the given id into the incremented one
func CreateComment(comment models.Comment, sid int) (models.Comment, error) {

	cid++
	comment.CommentID = cid

	for i := 0; i < len(issues); i++ {
		if issues[i].ID == sid {
			issues[i].Comments = append(issues[i].Comments, comment)
			comment.OnIssue = issues[i].IssueContent
			comments = append(comments, comment)
			return comment, nil
		}
	}
	return comment, fmt.Errorf("There is no issue with the corresponding id")
}

// DeleteComment deletes the comments of an issue, when that issue is deleted
func DeleteComment(cid int) error {
	for i := 0; i < len(comments); i++ {
		if comments[i].CommentID == cid {
			comments = append(comments[:i], comments[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("There is no comment with id %d to delete", cid)
}

// GetAllComments returns all comments
func GetAllComments() models.Comments {
	return comments
}

// GetCommentsByIssue returns all comments of an issue in []models.Comment
func GetCommentsByIssue(id int) []models.Comment {
	if len(issues) != 0 {
		for i := 0; i < len(issues); i++ {
			if issues[i].ID == id {
				return issues[i].Comments
			}
		}
	}
	return []models.Comment{}
}

// Initialize sample data
func init() {

	// issues
	CreateIssue(models.Issue{IssueContent: "Test create and edit functions", Status: "Pending"})
	CreateIssue(models.Issue{IssueContent: "Retreive data from datastorage", Status: "Pending"})
	CreateIssue(models.Issue{IssueContent: "Add Authentication", Status: "ToDo"})
	CreateIssue(models.Issue{IssueContent: "Add Docker", Status: "ToDo"})

	// comments
	CreateComment(models.Comment{Content: "This api needs some work"}, 4)
	CreateComment(models.Comment{Content: "This api is not public"}, 3)
	CreateComment(models.Comment{Content: "No authentication!? UNSAFE"}, 3)
}
