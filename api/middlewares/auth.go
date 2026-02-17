package middlewares

import (
	"context"
	"net/http"

	"go-api-template/api/helpers"
	"go-api-template/auth"
	db "go-api-template/database"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	ID  int   `json:"id"`
	Iat int64 `json:"iat"`
	jwt.RegisteredClaims
}

func JWTAuthMiddleware(secretKey []byte, store db.Storage) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			tokenString := r.Header.Get("Authorization")
			if tokenString == "" {
				errValue := helpers.ErrorMap[helpers.ErrUnauthorized]
				helpers.WriteJSON(w, http.StatusUnauthorized, nil, &errValue, "")
				return
			}

			claims, err := auth.ValidateJWT(secretKey, tokenString, store)
			if err != nil {

				helpers.WriteJSON(w, http.StatusUnauthorized, nil, &helpers.APIError{
					Code:    "UNAUTHORIZED",
					Message: err.Error(),
				}, "")
				return
			}

			ctx := context.WithValue(r.Context(), "userClaims", claims)
			next.ServeHTTP(w, r.WithContext(ctx))
		}
	}
}
