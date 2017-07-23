package models

type Issue struct {
	ID           int       `json:"id"`
	IssueContent string    `json:"issuecontent"`
	Status       string    `json:"status"`
	Comments     []Comment `json:"comments"`
}

type Issues []Issue
