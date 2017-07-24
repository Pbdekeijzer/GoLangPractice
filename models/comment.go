package models

type Comment struct {
	CommentID int    `json:"cid"`
	Content   string `json:"content"`
	OnIssue   int    `json:"issue"`
}

type Comments []Comment
