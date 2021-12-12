package service

import (
	"database/sql"
	"message-board/dao"
	"message-board/models"
)

func IsRepeatUsername(username string) (bool, error) {

	_, err := dao.SelectUsername(username)

	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func Register(user models.User) error {
	err := dao.CreateUser(user)
	return err
}

func Login(username, password string) (bool, error) {
	user, err := dao.Login(username)

	if err != nil {
		if password == user.Password {
			return true, err
		}
		return false, nil
	} else {
		if password == user.Password {
			return true, err
		}
		return false, nil
	}
}

func ReturnSecurityquestion(username string) (string, error) {
	securityquestion, err := dao.ReturnSecurityQuestion(username)
	if err != nil {
		return securityquestion, err
	}
	return securityquestion, nil
}

func IsSecurityanswerCorrect(username string) (string, error) {
	securityanswer, err := dao.SelectSecurityanswer(username)
	if err != nil {
		return securityanswer, err
	}
	return securityanswer, nil
}

func RetrievePassword(username string) (string, error) {
	oldpassword, error := dao.RetrievePassword(username)

	if error != nil {
		return oldpassword, error
	}
	return oldpassword, nil
}

func ChangePassword(username, newpassword string) error {
	err := dao.ChangePassword(username, newpassword)
	if err != nil {
		return err
	}
	return nil
}
