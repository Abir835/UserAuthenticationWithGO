package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"user-authentication-with-go/pkg/controllers"
)

func IsAuthorized(role string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("token")
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenStr := cookie.Value
			claims := &controllers.Claims{}

			token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
				return controllers.JwtKey, nil
			})
			if err != nil || !token.Valid || claims.Role != role {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
