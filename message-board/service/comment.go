package service

import (
	"message-board/dao"
	"message-board/models"
)

func AddComment(comment models.Comment) error {
	err := dao.InsertComment(comment)
	return err
}

func CancelComment(txt string) error {
	err := dao.DropComment(txt)
	return err
}
