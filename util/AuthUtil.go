package util

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

func GenerateJWTToken(userID int, username string) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24).Unix(), // Token expiration time (e.g., 24 hours)
		"iat":      time.Now().Unix(),                     // Token issuance time
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	secretKey := []byte("kaosdkaosdjiqwjeojo809812038123098#@#$$asdjoi")
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return "Bearer " + tokenString, nil
}
