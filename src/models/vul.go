package models

type Vul struct {
	ID  string `json:"id"`
	Url string `json:"url"`

	TaskID    string    `json:"task_id"`
	Task      *Task     `json:"task"`

	VulInfoID string    `json:"vul_info_id"`
	VulInfo   VulInfo   `json:"vul_info"`


	BonusInfo BonusInfo `json:"bonus_info"`
}
