package handlers

import (
    "encoding/json"
    "net/http"
    "go_api_vue_ws_v1/models"
    "github.com/gorilla/mux"
    "gorm.io/gorm"
    "golang.org/x/crypto/bcrypt"
)

var db *gorm.DB

// HashPassword hashes a password using bcrypt
func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
    return string(bytes), err
}

// CheckPasswordHash checks if the password is correct
func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}

// RegisterUser handles the registration of a new user
func RegisterUser(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Name            string `json:"name"`
        Email           string `json:"email"`
        Password        string `json:"password"`
        ConfirmPassword string `json:"confirmPassword"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    if request.Password != request.ConfirmPassword {
        http.Error(w, "Passwords do not match", http.StatusBadRequest)
        return
    }

    passwordHash, err := HashPassword(request.Password)
    if err != nil {
        http.Error(w, "Error hashing password", http.StatusInternalServerError)
        return
    }

    user := models.User{Name: request.Name, Email: request.Email, PasswordHash: passwordHash}
    db.Create(&user)

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "User created successfully"})
}

// LoginUser handles user login
func LoginUser(w http.ResponseWriter, r *http.Request) {
    var request struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }

    if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
        http.Error(w, "Invalid input", http.StatusBadRequest)
        return
    }

    var user models.User
    db.Where("email = ?", request.Email).First(&user)

    if user.ID == 0 || !CheckPasswordHash(request.Password, user.PasswordHash) {
        http.Error(w, "Invalid email or password", http.StatusUnauthorized)
        return
    }

    token, err := GenerateJWT(user.ID, user.Role)
    if err != nil {
        http.Error(w, "Failed to generate token", http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]string{"token": token})
}
