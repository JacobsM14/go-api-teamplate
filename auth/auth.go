package auth

import (
	"fmt"
	"os"
	"time"

	db "go-api-template/database"

	"github.com/golang-jwt/jwt/v5"
)

//var secretKey = []byte(os.Getenv("SECRET"))

type Claims struct {
	ID     int `json:"id"`
	IDRank int `json:"id_rank"`
	//DisplayName string `json:"display_name"`
	Iat int64 `json:"iat"`
	jwt.RegisteredClaims
}

func GenerateJWT(secretKey []byte, idUser int, idRank int, issuedAt time.Time) (string, error) {
	expirationTime := issuedAt.Add(18 * time.Hour)
	claims := Claims{
		ID:     idUser,
		IDRank: idRank,
		//DisplayName: displayName,
		Iat: issuedAt.Unix(),
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    os.Getenv("ISSUER"),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(issuedAt),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		return "", fmt.Errorf("Error signing token: %v", err)
	}

	return signedToken, nil
}

func ValidateJWT(secretKey []byte, tokenString string, store db.Storage) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", t.Header["alg"])
		}
		return secretKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("Failed to parse token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		return nil, fmt.Errorf("Invalid token claims")
	}

	lastTokenIssued, err := store.GetTokenIssuedDate(claims.ID)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve last token issued date for user ID %d: %v", claims.ID, err)
	}

	if time.Unix(claims.Iat, 0).Before(lastTokenIssued) {
		return nil, fmt.Errorf("Token is outdated, please log in again")
	}

	return claims, nil
}
