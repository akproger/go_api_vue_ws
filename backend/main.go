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

    "go_api_vue_ws_v1/handlers"
    "go_api_vue_ws_v1/models"
)

func main() {
    dsn := fmt.Sprintf("host=db user=%s password=%s dbname=%s port=5432 sslmode=disable",
        os.Getenv("POSTGRES_USER"), os.Getenv("POSTGRES_PASSWORD"), os.Getenv("POSTGRES_DB"))
    log.Println(dsn)

    db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
    if err != nil {
        log.Fatal("Не удалось подключиться к базе данных:", err)
    }

    db.AutoMigrate(&models.User{})
    handlers.SetDB(db)

    r := mux.NewRouter()
    handlers.InitializeRoutes(r)

    corsHandler := cors.New(cors.Options{
        AllowedOrigins: []string{"http://localhost:8082"},
        AllowedMethods: []string{"GET", "POST", "PUT", "DELETE"},
        AllowedHeaders: []string{"Content-Type"},
    }).Handler(r)

    log.Println("Сервер запущен на http://localhost:8080")
    log.Fatal(http.ListenAndServe(":8080", corsHandler))
}
