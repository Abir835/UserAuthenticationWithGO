package models

import "github.com/jinzhu/gorm"

type Purchase struct {
	gorm.Model
	BookId     int     `json:"bookId"`
	UserId     int     `json:"userId"`
	PCS        int     `json:"pcs"`
	Price      float32 `json:"price"`
	TotalPrice float32 `json:"totalPrice"`
}
