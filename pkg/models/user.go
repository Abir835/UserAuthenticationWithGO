package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
	Role     string // admin or normal
}
