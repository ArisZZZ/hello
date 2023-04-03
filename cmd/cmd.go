package cmd

import (
	"github.com/ariszzz/hello/internal/database"
	"github.com/ariszzz/hello/internal/router"
)

func RunServer(){
	database.Connect()
	r := router.New()
	r.Run(":8080")
}