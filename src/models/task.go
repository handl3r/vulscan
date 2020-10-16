package models

import "time"

type Task struct {
	ID        string    `json:"id"`
	Success   bool      `json:"success"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
}
