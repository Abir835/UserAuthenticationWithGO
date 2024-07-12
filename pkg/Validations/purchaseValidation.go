package Validations

import (
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
	"user-authentication-with-go/pkg/models"
)

func PurchaseValidationHandler(purchases []models.Purchase, db *gorm.DB) error {
	for _, purchase := range purchases {
		var book models.Book
		if err := db.First(&book, purchase.BookId).Error; err != nil {
			return fmt.Errorf("failed to find book: %v", err)
		}

		if purchase.Price != book.Price {
			return errors.New("price does not match")
		}

		totalPrice := purchase.Price * float32(purchase.PCS)
		if totalPrice != purchase.TotalPrice {
			return errors.New("total price does not match calculated price")
		}

		if purchase.PCS > book.PCS {
			return errors.New("purchase quantity exceeds stock")
		}
	}

	return nil
}
