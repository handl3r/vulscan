package clients

import (
	"encoding/json"
	"fmt"
	"github.com/valyala/fasthttp"
	"strings"
	"vulscan/src/enums"
)

type SqlmapClient struct {
	token    string
	endPoint string
	port     string
}

func NewSqlmapClient(token string, endPoint string, port string) *SqlmapClient {
	return &SqlmapClient{token: token, endPoint: endPoint, port: port}
}

func (smc *SqlmapClient) NewTask() (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	requestURL := fmt.Sprintf("%s:%s/task/new", smc.endPoint, smc.port)
	req.SetRequestURI(requestURL)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return "", enums.ErrClientNewTask
	}
	newTaskResponse := &NewTaskResponse{}
	err = json.Unmarshal(resp.Body(), newTaskResponse)
	if err != nil {
		return "", err
	}
	if !newTaskResponse.Success {
		return "", enums.ErrClientNewTask
	}
	return newTaskResponse.TaskID, nil
}

func (smc *SqlmapClient) SetOptionGET(taskID string) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	requestURL := fmt.Sprintf("%s:%s/option/%s/set", smc.endPoint, smc.port, taskID)
	req.SetRequestURI(requestURL)
	req.Header.SetMethod("POST")
	req.Header.Set("Content-Type", "application/json")
	optionGETBody := `{
    						"timeSec": 10,
    						"threads": 3
						}`
	req.SetBodyString(optionGETBody)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return enums.ErrClientSetOptionGET
	}
	status := &Status{}
	err = json.Unmarshal(resp.Body(), status)
	if !status.Success {
		return enums.ErrClientSetOptionGET
	}
	return nil
}

func (smc *SqlmapClient) SetOptionForPOST(taskID string, params map[string]string) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	stmtParams := make([]string, 0)
	for key, value := range params {
		if value == "" {
			stmtParams = append(stmtParams, key+"=1")
		} else {
			stmtParams = append(stmtParams, key+"="+value)
		}
	}
	dataPOST := strings.Join(stmtParams, "&")
	optionPOSTBody := fmt.Sprintf(`{"method": "POST", "timeSec": 10, "threads": 5, "data": "%s"}`, dataPOST)
	requestURL := fmt.Sprintf("%s:%s/option/%s/set", smc.endPoint, smc.port, taskID)
	req.Header.Set("Content-Type", "application/json")
	req.SetRequestURI(requestURL)
	req.Header.SetMethod("POST")
	req.SetBodyString(optionPOSTBody)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return enums.ErrClientSetOptionPOST
	}
	status := &Status{}
	err = json.Unmarshal(resp.Body(), status)
	if !status.Success {
		return enums.ErrClientSetOptionPOST
	}
	return nil
}

func (smc *SqlmapClient) StartScan(taskID, url string) error {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.Set("Content-Type", "application/json")
	requestURL := fmt.Sprintf("%s:%s/scan/%s/start", smc.endPoint, smc.port, taskID)
	req.SetRequestURI(requestURL)
	req.Header.SetMethod("POST")
	data := fmt.Sprintf(`{"url": "%s"}`, url)
	req.SetBodyString(data)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return err
	}
	scanResponse := &StartScanResponse{}
	err = json.Unmarshal(resp.Body(), scanResponse)
	if scanResponse.Success != true {
		return enums.ErrClientStartScan
	}
	return nil
}

func (smc *SqlmapClient) CheckTaskStatus(taskID string) (string, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	requestURL := fmt.Sprintf("%s:%s/scan/%s/status", smc.endPoint, smc.port, taskID)
	req.SetRequestURI(requestURL)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return "", err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return "", enums.ErrClientNewTask
	}
	statusScanResponse := &StatusScanResponse{}
	err = json.Unmarshal(resp.Body(), statusScanResponse)
	if err != nil {
		return "", err
	}
	if !statusScanResponse.Success || statusScanResponse.ReturnCode == enums.GeneralErrorOccurred ||
		statusScanResponse.ReturnCode == enums.UnhandledException {
		return "", enums.ErrClientEndScan
	}
	if statusScanResponse.Success {
		return statusScanResponse.Status, nil
	}
	return statusScanResponse.Status, nil
}

func (smc *SqlmapClient) GetData(taskID string) (*ResultScanData, int, error) {
	req := fasthttp.AcquireRequest()
	defer fasthttp.ReleaseRequest(req)
	req.Header.SetMethod("GET")
	requestURL := fmt.Sprintf("%s:%s/scan/%s/data", smc.endPoint, smc.port, taskID)
	req.SetRequestURI(requestURL)
	resp := fasthttp.AcquireResponse()
	defer fasthttp.ReleaseResponse(resp)
	err := fasthttp.Do(req, resp)
	if err != nil {
		return nil, 0, err
	}
	if resp.StatusCode() != fasthttp.StatusOK {
		return nil, 0, enums.ErrClientErrorGetResult
	}
	resultScanData := &ResultScanData{}
	err = json.Unmarshal(resp.Body(), resultScanData)
	if err != nil {
		return nil, 0, err
	}
	if len(resultScanData.Data) > 0 {
		for _, v := range resultScanData.Data {
			if v.Type == 1 {
				return resultScanData, enums.ResultExistVul, nil
			}
		}
	}
	return resultScanData, enums.ResultNotExistVul, nil
}
