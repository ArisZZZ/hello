package database

import (
	"fmt"
	"log"

	"github.com/ariszzz/hello/internal/entity"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "mangosteen"
	password = "123456"
	dbname   = "mangosteen_dev"
)


func Connect() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		log.Fatalln(err)
	}
	err = db.AutoMigrate(&entity.User{})
	DB = db
	if err != nil {
		panic("failed to migrate table")
	}
	log.Println("Connect db success")
}
