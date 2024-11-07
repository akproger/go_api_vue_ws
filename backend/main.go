package main

import (
    "fmt"
    "log"
    "net/http"
    "os"

    "github.com/gorilla/mux"
    "github.com/rs/cors"
    "gorm.io/driver/postgres"
    "gorm.io/gorm"

    "go_api_vue_ws_v1/handlers" // Путь к пакету handlers
    "go_api_vue_ws_v1/models"   // Путь к пакету models
)

var db *gorm.DB

func init() {
    var err error
    dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
    log.Println("Connecting to database with DSN:", dsn)

    db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatalf("Failed to connect to database: %v", err)
    }

    // Автоматическая миграция модели User
    if err := db.AutoMigrate(&models.User{}); err != nil {
        log.Fatalf("Failed to migrate database: %v", err)
    }
}

func main() {
    r := mux.NewRouter()

    // Передаем db в handlers, чтобы избежать nil значений
    handlers.SetupRoutes(r, db)

    r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("pong"))
    }).Methods("GET")

    // Обработчик для 404 ошибок
    r.NotFoundHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusNotFound)
        w.Write([]byte(`{"status": "error", "message": "Route not found"}`))
    })

    // Настройка CORS
    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8082"}, // Убедитесь, что это URL вашего фронтенд-приложения
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
    }).Handler(r)

    log.Println("Сервер запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
