package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Id       int    `json:"ID" gorm:"primary_key"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password int    `json:"password"`
}
