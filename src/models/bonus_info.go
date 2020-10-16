package models

type BonusInfo struct {
	ID    string `json:"id"`
	VulID string `json:"vul_id"`
	Data  interface{} `json:"data"`
}
