package models

// Project can create, read, update, delete by user owner
type Project struct {
	ID string `json:"id" gorm:"primary_key, not null"`
	VulNumber uint `json:"vul_number"`
	Domain string `json:"domain"`

	UserID string `json:"user_id"`
	User *User `json:"user"`

	Tasks []*Task `json:"tasks"`
}
