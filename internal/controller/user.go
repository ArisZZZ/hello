package controller

import (
	"fmt"

	"net/http"

	"github.com/ariszzz/hello/internal/database"
	"github.com/ariszzz/hello/internal/entity"
	"github.com/gin-gonic/gin"
)


func GetUserList(c *gin.Context) {
	users := []entity.User{}
	database.DB.Find(&users)

	c.JSON(http.StatusOK,users)
}

func CreateUser(c *gin.Context) {
	user := entity.User{}
	// 创建一个新用户
	c.Bind(&user)
	fmt.Println(user)
	result := database.DB.Create(&user)
	if result.Error != nil {
		panic("failed to create user")
	}
	fmt.Printf("Created user: %+v\n", user)
	c.JSON(http.StatusOK,"")
}
