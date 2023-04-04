package entity

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Uid   uint `json:"uid"`
	Uname string`json:"uname"`
	Email string`json:"email"`
}