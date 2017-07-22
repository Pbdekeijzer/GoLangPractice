package models

type Comment struct {
	CommentID uint   `json:"cid"`
	Content   string `json:"content"`
	OnIssue   *Issue `json"issue"`
}

type Comments []Comment
