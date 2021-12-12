package tool

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RespWrongPassword(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"密码错误，请回答密保问题以获取旧密码并修改密码",
	})
}

func Success(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"success",
	})
}

func Failure(ctx *gin.Context)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"failure",
	})
}

func ReturnQuestion(ctx *gin.Context,securityquestion string)  {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"securityquestion is: "+securityquestion,
	})
}

func ReturnPassword(ctx *gin.Context,oldpassword string) {
	ctx.JSON(http.StatusOK,gin.H{
		"info":"旧密码: "+ oldpassword,
	})
}