package models

// Project can create, read, update, delete by user owner. Has many segments
type Project struct {
	ID        string `json:"id" gorm:"primary_key, not null"`
	VulNumber uint   `json:"vul_number"`
	Name      string `json:"name"`
	Domain    string `json:"domain" gorm:"not null"`

	UserID string `json:"user_id" gorm:"not null"`
	User   *User  `json:"user"`

	Segments []Segment `json:"segments"`
}
