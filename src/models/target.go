package models

import (
	"net/url"
	"strings"
)

type Target struct {
	ID        string   `json:"id" gorm:"primary_key;not null"`
	VulID     string   `json:"vul_id"`
	URL       *url.URL `json:"url" gorm:"-"`
	RawURL    string   `json:"raw_url"`
	Method    string   `json:"method"`
	Params    string   `json:"params"` // separate by '*|*' ex: ["p1", "p2", "p3=2"]
	SegmentID string   `json:"segment_id"`
	Segment   *Segment `json:"segment" gorm:"-"`
}

func (t Target) GetMapParams() map[string]string {
	params := strings.Split(t.Params, "*|*")
	mapParams := make(map[string]string)
	for _, p := range params {
		temp := strings.Split(p, "=")
		if len(temp) == 1 {
			mapParams[temp[0]] = ""
		}
		if len(temp) == 2 {
			mapParams[temp[0]] = temp[1]
		}
	}
	return mapParams
}
