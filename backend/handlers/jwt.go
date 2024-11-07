package handlers

import (
    "github.com/golang-jwt/jwt/v4"
    "time"
)

var jwtKey = []byte("your_secret_key")

func GenerateJWT(userID uint, role string) (string, error) {
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
        "userID": userID,
        "role":   role,
        "exp":    time.Now().Add(24 * time.Hour).Unix(),
    })
    return token.SignedString(jwtKey)
}
