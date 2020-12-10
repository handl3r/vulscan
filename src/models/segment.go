package models

import "time"

// Segment can create, read , delete. Based on time of scanning. Belong to one Project. Has many Targets and Vuls
type Segment struct {
	ID           string    `json:"id" gorm:"primary_key, not null"`
	IsScanned    bool      `json:"is_scanned" gorm:"default:false"` // to check if user request scan a scanned segment -> reject
	CreatedAt    time.Time `json:"created_at" gorm:"default:CURRENT_TIMESTAMP"`
	DeletedAt    time.Time `json:"deleted_at" gorm:"default:CURRENT_TIMESTAMP"`
	TargetNumber uint      `json:"target_number"`
	VulNumber    uint      `json:"vul_number"`
	Targets      []Target  `json:"targets"`

	ProjectID string   `json:"project_id" gorm:"not null"`
	Project   *Project `json:"project"`

	UserID string `json:"user_id"`

	Vuls []Vul `json:"vuls"`
}
