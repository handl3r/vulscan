package models

import (
	"time"
)

type User struct {
	ID        string    `json:"id" gorm:"primary_key"`
	Email     string    `json:"email" gorm:"not null"`
	Username  string    `json:"username" gorm:"not null"`
	Password  string    `json:"-" gorm:"not null"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`

	Projects []Project `json:"projects"`
}
