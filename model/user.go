package model

type User struct {
	ID        uint   `json:"id" gorm:"primaryKey"`
	UserID    string `json:"user_id" gorm:"unique,not null"`
	Name      string `json:"name" gorm:"not null"`
	Email     string `json:"email" gorm:"unique;not null"`
	Password  string `json:"password" gorm:"not null"`
	CreatedAt int64  `json:"created_at" gorm:"not null"`
}
