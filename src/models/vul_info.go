package models

type VulInfo struct {
	ID            string      `json:"id" gorm:"primary_key, not null"`
	Success       bool        `json:"success"`
	Data          interface{} `json:"data"`
	ErrorResponse interface{} `json:"error"`
}
