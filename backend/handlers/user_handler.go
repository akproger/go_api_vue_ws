package handlers

import (
    "encoding/json"
    "net/http"
    "go_api_vue_ws_v1/models"
    "github.com/gorilla/mux"
)

func GetUsers(w http.ResponseWriter, r *http.Request) {
    var users []models.User
    result := db.Find(&users)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
    var user models.User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    result := db.Create(&user)
    if result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    if result := db.First(&user, params["id"]); result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    if result := db.First(&user, params["id"]); result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    var updatedUser models.User
    if err := json.NewDecoder(r.Body).Decode(&updatedUser); err != nil {
        http.Error(w, "Invalid request payload", http.StatusBadRequest)
        return
    }

    user.Name = updatedUser.Name
    user.Email = updatedUser.Email

    if result := db.Save(&user); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(user)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var user models.User
    if result := db.First(&user, params["id"]); result.Error != nil {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    if result := db.Delete(&user); result.Error != nil {
        http.Error(w, result.Error.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusNoContent)
}
