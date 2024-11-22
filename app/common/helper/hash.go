package common_helper

import (
	"crypto/rand"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/phucpham-infinity/go-nextpress/app/context"
)

func GenerateActivationKey(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func GenerateJWT(userID string, email string, username string) (string, error) {
	jwtConfig := context.AppContext().GetConfig().JWT

	secretKey := jwtConfig.Secret
	expirationDays := jwtConfig.ExpirationDays
	expirationTime := time.Now().Add(time.Duration(expirationDays) * 24 * time.Hour)

	claims := jwt.MapClaims{
		"id":       userID,
		"email":    email,
		"username": username,
		"exp":      expirationTime.Unix(),
		"iat":      time.Now().Unix(),
		"iss":      "be-gopress",
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
