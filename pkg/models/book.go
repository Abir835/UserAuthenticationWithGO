package models

import "github.com/jinzhu/gorm"

type Book struct {
	gorm.Model
	Name        string  `gorm:"unique" json:"name"`
	Author      string  `json:"author"`
	Publication string  `json:"publication"`
	PCS         int     `json:"pcs"`
	Price       float32 `json:"price"`
}
