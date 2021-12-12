package api

import "github.com/gin-gonic/gin"

func InitEngine()  {
	engine := gin.Default()

	engine.POST("/register",register)
	engine.POST("/login",login)
	engine.GET("returnquestion",returnquestion)
	engine.POST("/changepassword",changepassword)

	engine.POST("comment",Comment)
	engine.POST("cancelcomment",CancelComment)

	engine.Run(":8090")
}
