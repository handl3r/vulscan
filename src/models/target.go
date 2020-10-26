package models

type Target struct {
	ID        string `json:"id" gorm:"primary key, not null"`
	VulID     string `json:"vul_id"`
	URL       string `json:"url"`
	Method    int    `json:"method"`
	Params    string `json:"params"` // separate by ','
	SegmentID string `json:"segment_id"`
	Segment   *Segment
}
