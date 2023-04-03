package cmd

import "github.com/ariszzz/hello/internal/database"

func RunServer(){
	database.Connect()
}