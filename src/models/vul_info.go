package models

// VulInfo describe information about a Vul. Include Title (type exploit),... Belong and depend to a Vul
type VulInfo struct {
	Vul             Vul         `json:"vul"`
	ID              string      `json:"id" gorm:"primary_key, not null"`
	VulID           string      `json:"vul_id"`
	Title           string      `json:"title"`
	Payload         string      `json:"payload"`
	Vector          string      `json:"vector"`
	TemplatePayload string      `json:"template_payload"`
	ErrorResponse   interface{} `json:"error"`
	TrueCode        uint        `json:"true_code"`
	FalseCode       uint        `json:"false_code"`
	MatchRatio      float32     `json:"match_ratio"`
	Success         bool        `json:"success"`
}
