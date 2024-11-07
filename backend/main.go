package main

import (
    "fmt"
    "log"
    "net/http"
    "os"
    "context"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "go_api_vue_ws_v1/handlers"
    "go_api_vue_ws_v1/models"
)

var db *gorm.DB

func init() {
    var err error
    dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
    log.Println(dsn)

    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }
}

func main() {
    dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
    log.Println(dsn)

    var err error
    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных:", err)
    }

    db.AutoMigrate(&models.User{})

    r := mux.NewRouter()
    r.HandleFunc("/register", handlers.RegisterUser).Methods("POST")
    r.HandleFunc("/login", handlers.LoginUser).Methods("POST")

    protected := r.PathPrefix("/admin").Subrouter()
    protected.Use(Authenticate, AuthorizeAdmin)
    protected.HandleFunc("/users", handlers.GetUsers).Methods("GET")

    r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    }).Methods("GET")

    r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"status": "error", "message": "Route not found"}`))
    })

    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8082"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type", "Authorization"},
    }).Handler(r)

    log.Println("Server is running at http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
