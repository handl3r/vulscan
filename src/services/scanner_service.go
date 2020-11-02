package services

import (
	"log"
	"sync"
	"time"
	"vulscan/src/clients"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/repositories"
)

type ScannerService struct {
	segmentRepository repositories.SegmentRepository
	sqlmapClient      clients.SqlmapClient
}

func (ss *ScannerService) ScanMultiTargets(targets []models.Target) {

	resultChan := make(chan *models.Vul)
	var wg sync.WaitGroup
	for i := range targets {
		wg.Add(1)
		go func(target models.Target) {
			_ = ss.ScanTarget(target, resultChan, &wg)
		}(targets[i])
	}
	wg.Wait()
	close(resultChan)
}

func (ss *ScannerService) ScanTarget(target models.Target, resultChan chan *models.Vul, wg *sync.WaitGroup) error {
	defer wg.Done()
	var err error
	taskID, err := ss.sqlmapClient.NewTask()
	if err != nil {
		log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
		return enums.ErrSystem
	}
	switch target.Method {
	case enums.GET:
		err = ss.sqlmapClient.SetOptionGET(taskID)
	case enums.POST:
		mapParams := target.GetMapParams()
		err = ss.sqlmapClient.SetOptionForPOST(taskID, mapParams)
	}
	if err != nil {
		log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
		return enums.ErrSystem
	}
	err = ss.sqlmapClient.StartScan(taskID, target.URL)
	if err != nil {
		log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
		return enums.ErrSystem
	}
	// decrease wait duration each time check status
	for true {
		status, err := ss.sqlmapClient.CheckTaskStatus(taskID)
		if err != nil {
			log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
			return err
		}
		if status == enums.StatusRunning {
			time.Sleep(5 * time.Second)
		}
		if status == enums.StatusTerminated {
			break
		}
	}
	_, isVul, err := ss.sqlmapClient.GetData(taskID)
	if err != nil {
		log.Printf("Error when get result of task [TargetID: %s] [TaskID: %s] [Error: %s]", target.ID, taskID, err)
		return err
	}
	if isVul == enums.ResultExistVul {
		vul := models.NewVulWithTarget(target)
		resultChan <- vul
	}
	return nil
}
