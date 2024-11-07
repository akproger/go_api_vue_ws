package main

import (
    "context"
    "net/http"
    "github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

func Authenticate(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        tokenStr := r.Header.Get("Authorization")
        if tokenStr == "" {
            http.Error(w, "Missing token", http.StatusUnauthorized)
            return
        }

        claims := &jwt.MapClaims{}
        token, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
            return jwtKey, nil
        })

        if err != nil || !token.Valid {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        r = r.WithContext(context.WithValue(r.Context, "user", claims))
        next.ServeHTTP(w, r)
    })
}

func AuthorizeAdmin(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        user := r.Context().Value("user").(*jwt.MapClaims)
        if role, ok := (*user)["role"].(string); !ok || role != "admin" {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }
        next.ServeHTTP(w, r)
    })
}
