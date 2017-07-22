package models

import "time"

type Issue struct {
	ID           uint       `json:"id"`
	DateCreated  time.Time  `json:"createdOn"`
	DateModified time.Time  `json:"modifiedOn"`
	IssueContent string     `json:"issuecontent"`
	Status       string     `json:"status"`
	Comments     *[]Comment `json:"comments"`
}

type Issues []Issue
