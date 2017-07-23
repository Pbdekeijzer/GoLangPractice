package datastorage

import (
	"fmt"

	"encoding/json"

	"github.com/pbdekeijzer/GoLangPractice/models"
)

//Issue
var issues models.Issues
var sid int

func CreateIssue(issue models.Issue) {
	sid++
	issue.ID = sid
	issues = append(issues, issue)
}

func DeleteIssue(id int) error {

	for i, s := range issues {
		if s.ID == id {

			//delete all the comments from the issue
			if len(issues[i].Comments) != 0 {
				for _, comment := range issues[i].Comments {
					deleteComment(comment.CommentID)
				}
			}
			//delete the issue itselfs
			issues = append(issues[:i], issues[i+1:]...)
			return nil
		}
	}
	fmt.Println("test")
	return fmt.Errorf("There is no issue with id %d to delete", id)
}

func EditIssue(issue models.Issue) error {
	for i := 0; i < len(issues); i++ {
		if issues[i].ID == issue.ID {
			issues[i] = issue
		}
	}
	return fmt.Errorf("Not able to edit issue with id: %d", issue.ID)
}

func GetAllIssues() models.Issues {
	return issues
}

func GetIssue(id int) models.Issue {
	for i := 0; i < len(issues); i++ {
		if issues[i].ID == id {
			return issues[i]
		}
	}
	//else return empty Issue
	return models.Issue{}
}

//Comment
var comments models.Comments
var cid int

//CreateComment and appends the comment to the comments of the corresponding issue
func CreateComment(comment models.Comment, sid int) {
	cid++
	comment.CommentID = cid

	for i := 0; i < len(issues); i++ {
		if issues[i].ID == sid {
			issues[i].Comments = append(issues[i].Comments, comment)
			comment.OnIssue = issues[i].IssueContent
			comments = append(comments, comment)
			return
		}
	}
}

//deleteComment deletes the comments of an issue, when that issue is deleted
func deleteComment(cid int) error {
	for i := 0; i < len(comments); i++ {
		if comments[i].CommentID == cid {
			comments = append(comments[:i], comments[i+1:]...)
			return nil
		}
	}
	return fmt.Errorf("There is no comment with id %d to delete", cid)
}

//GetAllComments returns all comments
func GetAllComments() models.Comments {
	return comments
}

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

//add some test data
func init() {
	//add issues to the issue slice
	CreateIssue(models.Issue{IssueContent: "Test create and edit functions", Status: "Pending"})
	CreateIssue(models.Issue{IssueContent: "Retreive data from datastorage", Status: "Pending"})
	CreateIssue(models.Issue{IssueContent: "Add Authentication", Status: "ToDo"})
	CreateIssue(models.Issue{IssueContent: "Add Docker", Status: "ToDo"})

	//test creation of comment
	CreateComment(models.Comment{Content: "This api needs some work"}, 4)
	CreateComment(models.Comment{Content: "This api is not public"}, 3)
	fmt.Println(issues[3].Comments[0].Content)

	//test deletion of issue
	fmt.Println("Current comments are:")
	for _, element := range GetAllComments() {
		fmt.Println(element)
	}

	issue := models.Issue{IssueContent: "TEST"}
	b, err := json.Marshal(issue)

	if err != nil {
		fmt.Println("sucks")
	}
	fmt.Println(string(b))

	DeleteIssue(4)
	fmt.Println("Comments after delete are:")
	for _, element := range GetAllComments() {
		fmt.Println(element)
	}

	fmt.Println(GetCommentsByIssue(3))

	//test edit issue: Test create and edit functions -> Succesfully tested create and edit functions
	EditIssue(models.Issue{ID: 1, IssueContent: "Succesfully tested create and edit functions", Status: "Completed"})
	fmt.Println(issues[0].IssueContent)

}
