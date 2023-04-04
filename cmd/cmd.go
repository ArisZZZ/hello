package cmd

import (
	"github.com/ariszzz/hello/internal/database"
	"github.com/ariszzz/hello/internal/router"
)

type User struct {
	uid        uint `gorm:"primaryKey"`
	uname      string
	email       string
	created_at string
	updated_at string
}

func RunServer(){
	database.Connect()
	r := router.New()
	r.Run(":8080")
}