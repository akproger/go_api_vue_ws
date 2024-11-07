package handlers

import (
    "github.com/gorilla/mux"
)

// InitializeRoutes инициализирует маршруты для приложения
func InitializeRoutes(r *mux.Router) {
    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/users", CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
    r.HandleFunc("/register", RegisterUser).Methods("POST")
    r.HandleFunc("/login", LoginUser).Methods("POST")
}
