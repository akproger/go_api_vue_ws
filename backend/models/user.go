package models

type User struct {
    ID           uint   `json:"id" gorm:"primaryKey"`
    Name         string `json:"name"`
    Email        string `json:"email" gorm:"unique"`
    PasswordHash string `json:"-"`
    Role         string `json:"role" gorm:"default:user"` // роль по умолчанию — "user"
}