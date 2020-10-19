package models

// Vul of a target. Can create, read, delete by delete segment, can not update. Belong to a Segment and has many VulInfo
// and BonusInfo
type Vul struct {
	ID        string `json:"task_id" gorm:"primary_key, not null"`
	Target    string `json:"target"`
	Method    string `json:"method"`
	Parameter string `json:"parameter"`
	Suffix    string `json:"suffix"`

	SegmentID string   `json:"segment_id"`
	Segment   *Segment `json:"segment"`

	VulInfo []VulInfo `json:"vul_info"`

	BonusInfo []string `json:"bonus_info"`
}
