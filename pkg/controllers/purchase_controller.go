package controllers

import (
	"encoding/json"
	"github.com/jinzhu/gorm"
	"net/http"
	"user-authentication-with-go/pkg/Validations"
	"user-authentication-with-go/pkg/models"
	"user-authentication-with-go/pkg/utils"
)

type PurchaseResponse struct {
	BookName   string  `json:"bookName"`
	PCS        int     `json:"pcs"`
	Price      float32 `json:"price"`
	TotalPrice float32 `json:"totalPrice"`
}

func PurchaseHandler(db *gorm.DB) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		claims, _ := utils.GetUserIdByToken(r)
		var purchases []models.Purchase

		if err := json.NewDecoder(r.Body).Decode(&purchases); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err := Validations.PurchaseValidationHandler(purchases, db)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var responses []PurchaseResponse

		for i := range purchases {

			purchases[i].UserId = claims.UserId

			var book models.Book
			if err := db.First(&book, purchases[i].BookId).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			if err := db.Create(&purchases[i]).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			book.PCS = book.PCS - purchases[i].PCS
			if err := db.Save(&book).Error; err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			responses = append(responses, PurchaseResponse{
				BookName:   book.Name,
				PCS:        purchases[i].PCS,
				Price:      purchases[i].Price,
				TotalPrice: purchases[i].TotalPrice,
			})
		}

		w.WriteHeader(http.StatusCreated)

		json.NewEncoder(w).Encode(responses)
	}
}
