package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/models"
	"message-board/service"
	"message-board/tool"
)

func register(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	securityquestion := ctx.PostForm("securityquestion")
	securityanswer := ctx.PostForm("securityanswer")

	user := models.User{
		Username:         username,
		Password:         password,
		Securityquestion: securityquestion,
		Securityanswer:   securityanswer,
	}

	flag, err := service.IsRepeatUsername(user.Username)

	if err != nil {
		fmt.Println(err)
	}

	if !flag && err == nil {
		err := service.Register(user)
		if err != nil {
			panic(err)
		}
		tool.Success(ctx)
	} else{
		tool.Failure(ctx)
	}
}

func login(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	user := models.User{
		Username:         username,
		Password:         password,
	}

	flag, err := service.Login(user.Username, user.Password)
	if err != nil {
		fmt.Println(err)
	}
	if flag != true {
		tool.RespWrongPassword(ctx)

	}else{
		tool.Success(ctx)
	}
}

//返回密保问题
func returnquestion(ctx *gin.Context)  {
	username := ctx.PostForm("username")

	securityquestion, err := service.ReturnSecurityquestion(username)
	if err != nil {
		fmt.Println(err)
	}

	tool.ReturnQuestion(ctx, securityquestion)
}

//重置密码
func changepassword(ctx *gin.Context) {
	username := ctx.PostForm("username")
	securityanswer := ctx.PostForm("securityanswer")
	newpassword := ctx.PostForm("newpassword")

	correctanswer, err := service.IsSecurityanswerCorrect(username)
	if err != nil {
		fmt.Println(err)
	}

	if correctanswer == securityanswer {
		oldpassword,err := service.RetrievePassword(username)
		if err != nil {
			fmt.Println(err)
		}

		tool.ReturnPassword(ctx,oldpassword)

		error := service.ChangePassword(username,newpassword)
		if error != nil{
			fmt.Println(error)
		}
		tool.Success(ctx)
	} else {
		tool.Failure(ctx)
	}
}
