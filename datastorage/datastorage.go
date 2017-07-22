package datastorage

import (
	"fmt"

	"github.com/pbdekeijzer/GoLangPractice/models"
)

var issues models.Issues
var id uint

func CreateIssue(issue models.Issue) {
	id++
	issue.ID = id
	issues = append(issues, issue)
}

func DeleteIssue(id uint) error {
	for i, s := range issues {
		if s.ID == id {
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

	return fmt.Errorf("There is no issue with id %d to edit")
}

func init() {
	CreateIssue(models.Issue{ID: 1})
	CreateIssue(models.Issue{ID: 2})
}
