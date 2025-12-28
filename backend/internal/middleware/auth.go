package middleware

import (
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	service "github.com/lanxre/kyokusulib/internal/services"
)

type contextKey string

const (
	UserIDKey    contextKey = "userID"
	UserEmailKey contextKey = "userEmail"
)

func AuthMiddleware(next http.HandlerFunc, jwtSecret string) http.HandlerFunc {
	secretKey := []byte(jwtSecret)

	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("KYOKUSU_API_TOKEN")
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		tokenString := cookie.Value

		claims := &service.Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
				return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
			}
			return secretKey, nil
		})

		if err != nil || !token.Valid {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UserEmailKey, claims.Email)
		ctx = context.WithValue(ctx, UserRoleKey, claims.Role) 
		
		next(w, r.WithContext(ctx))
	}
}
