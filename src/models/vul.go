package models

type Vul struct {
	TaskID  string `json:"task_id" gorm:"primary_key, not null"`
	Target string `json:"target"`

	SegmentID    string    `json:"segment_id"`
	Segment      *Segment     `json:"segment"`

	VulInfoID string    `json:"vul_info_id"`
	VulInfo   VulInfo   `json:"vul_info"`


	BonusInfo BonusInfo `json:"bonus_info"`
}
