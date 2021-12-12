package models

import "time"

type Comment struct {
	Username string
	Postman string
	Txt string
	CommentTime time.Time
}
