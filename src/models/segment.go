package models

import "time"

// Segment can create, read , delete. Based on time of scanning. Belong to one Project. Has many Targets and Vuls
type Segment struct {
	ID           string    `json:"id" gorm:"primary_key, not null"`
	CreatedAt    time.Time `json:"created_at"`
	DeletedAt    time.Time `json:"deleted_at"`
	TargetNumber uint      `json:"target_number"`
	VulNumber    uint      `json:"vul_number"`
	Targets      []string  `json:"targets"`

	ProjectID string   `json:"project_id"`
	Project   *Project `json:"project"`

	Vuls []Vul `json:"vuls"`
}
