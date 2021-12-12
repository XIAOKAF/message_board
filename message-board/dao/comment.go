package dao

import "message-board/models"

func InsertComment(comment models.Comment) error {
	_,err:=DB.Exec("insert into comment(username,post_username,txt,post_time)"+"values(?,?,?,?);",comment.Username,comment.Postman,comment.Txt,comment.CommentTime)
	return err
}

func DropComment(txt string) error {
	_,err := DB.Exec("delete from comment where txt =?",txt)
	return err
}
