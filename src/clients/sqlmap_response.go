package clients

type NewTaskResponse struct {
	TaskID  string `json:"taskid"`
	Success bool   `json:"success"`
}

type StartScanResponse struct {
	Engineid uint32 `json:"engineid"`
	Success  bool   `json:"success"`
}

type StatusScanResponse struct {
	Status     string `json:"status"`
	ReturnCode uint   `json:"returncode"`
	Success    bool   `json:"success"`
}

type ResultScanData struct {
	Success bool     `json:"success"`
	Data    []DataScan `json:"data"`
	Error   string   `json:"error"`
}

type DataScan struct {
	Status int         `json:"status"`
	Type   int         `json:"type"`
	Value  interface{} `json:"value"`
}
