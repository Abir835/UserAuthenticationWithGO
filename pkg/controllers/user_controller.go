package controllers

import (
	"net/http"
)

func AdminDashboard() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to the admin dashboard!"))
	}
}

func UserProfile() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Welcome to your profile!"))
	}
}
