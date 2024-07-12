package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"user-authentication-with-go/pkg/models"
)

type BookResponse struct {
	Name   string `json:"name"`
	Author string `json:"author"`
}

func CreateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book
		if err := json.NewDecoder(r.Body).Decode(&books); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		var responses []BookResponse
		for _, book := range books {
			db.Create(&book)
			responses = append(responses, BookResponse{Name: book.Name, Author: book.Author})
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(responses)
	}
}

func GetBooks(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var books []models.Book
		db.Find(&books)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(books)
	}
}

func GetBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var book models.Book
		db.First(&book, params["id"])
		if book.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Book not found")
			return
		}
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func UpdateBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var book models.Book
		db.First(&book, params["id"])
		if book.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Book not found")
			return
		}
		json.NewDecoder(r.Body).Decode(&book)
		db.Save(&book)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(book)
	}
}

func DeleteBook(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)
		var book models.Book
		db.First(&book, params["id"])
		if book.ID == 0 {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode("Book not found")
			return
		}
		db.Delete(&book)
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode("Book deleted")
	}
}
