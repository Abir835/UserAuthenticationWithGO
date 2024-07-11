package controllers

import (
	"encoding/json"
	"github.com/dgrijalva/jwt-go"
	"github.com/jinzhu/gorm"
	"net/http"
	"time"
	"user-authentication-with-go/pkg/models"
	"user-authentication-with-go/pkg/utils"
)

var JwtKey = []byte("my_secret_key")

type Claims struct {
	Email string `json:"email"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserResponse struct {
	Username string `json:"username"`
}

func Register(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user models.User
		if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		hashedPassword, err := utils.HashPassword(user.Password)
		if err != nil {
			http.Error(w, "Error hashing password", http.StatusInternalServerError)
			return
		}
		user.Password = hashedPassword
		user.OTPExpiry = time.Now()
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}

		utils.SendEmail(user.Email, "Registration Successful", "Welcome to Bookstore!")

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(map[string]string{"username": user.Username})
	}
}

func Login(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.Where("email = ?", creds.Email).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if !utils.CheckPasswordHash(creds.Password, user.Password) {
			http.Error(w, "Invalid credentials", http.StatusUnauthorized)
			return
		}

		otp := utils.GenerateOTP()
		user.OTP = otp
		user.OTPExpiry = time.Now().Add(1 * time.Minute)
		db.Save(&user)

		utils.SendEmail(user.Email, "OTP for Login", "Your OTP is: "+otp)

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "OTP sent to your email"})
	}
}

func VerifyOTP(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req struct {
			Email string `json:"email"`
			OTP   string `json:"otp"`
		}
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request payload", http.StatusBadRequest)
			return
		}

		var user models.User
		if err := db.Where("email = ?", req.Email).First(&user).Error; err != nil {
			http.Error(w, "User not found", http.StatusUnauthorized)
			return
		}

		if user.Email == "" || !utils.ValidateOTP(&user, req.OTP) {
			http.Error(w, "Invalid OTP", http.StatusUnauthorized)
			return
		}

		if user.OTPExpiry.Before(time.Now()) {
			http.Error(w, "OTP has expired", http.StatusUnauthorized)
			return
		}

		tokenString, err := utils.GenerateJWT(user.Email, user.Role)
		if err != nil {
			http.Error(w, "Error generating token", http.StatusInternalServerError)
			return
		}

		http.SetCookie(w, &http.Cookie{
			Name:    "token",
			Value:   tokenString,
			Expires: time.Now().Add(24 * time.Hour),
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Login Successfully"})
	}
}

func Logout() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.SetCookie(w, &http.Cookie{
			Name:     "token",
			Value:    "",
			Expires:  time.Now().Add(-1 * time.Hour), // Immediate expiration
			HttpOnly: true,
			Path:     "/",
		})

		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(map[string]string{"message": "Logout successful"})
	}
}
