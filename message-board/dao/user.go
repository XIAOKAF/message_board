package dao

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"message-board/models"
)

func SelectUsername(username string) (models.User, error) {
	user := models.User{}

	tx, _ := DB.Begin()

	row := DB.QueryRow("select id,username from Users where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
		fmt.Println(row.Err())
	}

	err := row.Scan(&user.Id, &user.Username)
	if err != nil {
		return user, err
	}
	tx.Commit()
	return user, nil
}

func CreateUser(user models.User) error {
	_, err := DB.Exec("INSERT INTO Users(username,password,securityquestion,securityanswer)"+"values (?,?,?,?);", user.Username, user.Password, user.Securityquestion, user.Securityanswer)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Login(username string) (models.User, error) {
	user := models.User{}

	row := DB.QueryRow("select password from Users where username=?", username)
	if row.Err() != nil {
		return user, row.Err()
		fmt.Println(row.Err())
	}

	err := row.Scan(&user.Password)
	if err != nil {
		return user, err
	}
	return user, nil
}

func ReturnSecurityQuestion(username string) (string,error)  {
	user := models.User{}

	row := DB.QueryRow("select securityquestion from Users where username=?", username)
	if row.Err() != nil {
		return user.Securityquestion, row.Err()
		fmt.Println(row.Err())
	}
	err := row.Scan(&user.Securityquestion)
	if err != nil {
		return user.Securityquestion, err
	}

	return user.Securityquestion, nil
}

func SelectSecurityanswer(username string) (string, error) {
	user := models.User{}

	row := DB.QueryRow("select securityanswer from Users where username=?", username)
	if row.Err() != nil {
		return user.Securityanswer, row.Err()
		fmt.Println(row.Err())
	}

	err := row.Scan(&user.Securityanswer)
	if err != nil {
		return user.Securityanswer, err
	}

	return user.Securityanswer, nil
}

func RetrievePassword(username string) (string, error) {
	user := models.User{}

	row:= DB.QueryRow("select password from Users where username=?", username)
	if row.Err() != nil {
		return user.Password, row.Err()
	}
//cuo
	err := row.Scan(&user.Password)

	if err != nil {
		return user.Password, err
	}

	return user.Password, nil
}

func ChangePassword(username, newpassword string) error {
	sqlStr := "update Users set password=? where username=?"
	ret,err := DB.Exec(sqlStr,newpassword,username)
	if err!=nil{
		return err
	}

	_,error :=ret.RowsAffected()
	if error !=nil{
		return error
	}
	return nil
}
