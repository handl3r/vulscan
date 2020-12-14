package models

// Vul of a target. Can create, read, delete by delete segment, can not update. Belong to a Segment and has many VulInfo
// and BonusInfo
type Vul struct {
	ID        string `json:"task_id" gorm:"primary_key;not null"`
	TargetID  string `json:"target_id" gorm:"not null"`
	Target    Target `json:"target" gorm:"-"`
	Method    string `json:"method"`
	Parameter string `json:"parameter"`
	Suffix    string `json:"suffix"`

	SegmentID string   `json:"segment_id"`
	Segment   *Segment `json:"segment" gorm:"-"`

	VulInfo []VulInfo `json:"vul_info" gorm:"-"`

	BonusInfo string `json:"bonus_info"`
}

func NewVulWithTarget(target Target) *Vul {
	return &Vul{
		TargetID: target.ID,
		Target:    target,
		Method:    target.Method,
		SegmentID: target.SegmentID,
	}
}
