package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"message-board/models"
	"message-board/service"
	"message-board/tool"
	"time"
)

func Comment(ctx *gin.Context)  {
	postusername := ctx.PostForm("postusername")
	username:=ctx.PostForm("username")
	txt := ctx.PostForm("txt")

	comment := models.Comment{
		Username: username,
		Postman: postusername,
		Txt: txt,
		CommentTime: time.Now(),
	}

	err := service.AddComment(comment)
	if err != nil{
		fmt.Println(err)
		tool.Failure(ctx)
	}else {
		tool.Success(ctx)
	}
}


func CancelComment(ctx *gin.Context)  {
	txt := ctx.PostForm("txt")

	err:=service.CancelComment(txt)
	if err != nil{
		fmt.Println(err)
		tool.Failure(ctx)
	}else {
		tool.Success(ctx)
	}
}
