package models

import "time"

type Segment struct {
	ID        string    `json:"id" gorm:"primary_key, not null"`
	CreatedAt time.Time `json:"created_at"`
	DeletedAt time.Time `json:"deleted_at"`
	VulNumber uint      `json:"vul_number"`

	Vul []Vul `json:"vul"`
}
