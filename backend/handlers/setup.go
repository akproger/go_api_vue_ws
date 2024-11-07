package handlers

import (
    "github.com/gorilla/mux"
    "gorm.io/gorm"
)

var db *gorm.DB

// SetupRoutes initializes routes and assigns the database connection.
func SetupRoutes(r *mux.Router, database *gorm.DB) {
    db = database // Assign the passed db connection to the handlers package variable

    r.HandleFunc("/users", GetUsers).Methods("GET")
    r.HandleFunc("/users", CreateUser).Methods("POST")
    r.HandleFunc("/users/{id}", GetUser).Methods("GET")
    r.HandleFunc("/users/{id}", UpdateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", DeleteUser).Methods("DELETE")
}
