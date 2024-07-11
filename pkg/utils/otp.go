package utils

import (
	"crypto/rand"
	"encoding/base32"
	"github.com/jinzhu/gorm"
	"time"
	"user-authentication-with-go/pkg/models"
)

func GenerateOTP() string {
	otp := make([]byte, 10)
	_, err := rand.Read(otp)
	if err != nil {
		return ""
	}
	return base32.StdEncoding.WithPadding(base32.NoPadding).EncodeToString(otp)[:6]
}

func StoreOTP(user *models.User, db *gorm.DB) (string, error) {
	otp := GenerateOTP()
	user.OTP = otp
	user.OTPExpiry = time.Now().Add(5 * time.Minute)
	db.Save(user)
	return otp, nil
}

func ValidateOTP(user *models.User, otp string) bool {
	if time.Now().After(user.OTPExpiry) {
		return false
	}
	return user.OTP == otp
}
