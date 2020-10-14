package models

import "time"

// Task is result of each scan time
type Task struct {
	ID      string `json:"id"`
	Success bool   `json:"success"`
	Status  string `json:"status"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`

	ProjectID string  `json:"project_id" gorm:"not null"`
	Project   *Project `json:"project"`

	VulList []*Vul `json:"vul_list"`
}
