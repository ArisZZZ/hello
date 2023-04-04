package router

import (
	"github.com/ariszzz/hello/internal/controller"
	"github.com/gin-gonic/gin"
)


func New() *gin.Engine{
	r := gin.Default()
	r.GET("/ping",controller.Ping)
	r.GET("/user/list",controller.GetUserList)
	r.POST("/user",controller.CreateUser)
	return r
}