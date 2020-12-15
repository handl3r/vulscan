package services

import (
	"log"
	"net/http"
	"sync"
	"time"
	"vulscan/src/adapter/clients"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/packages"
)

type ScannerService struct {
	segmentRepository repositories.SegmentRepository
	vulRepository     repositories.VulRepository
	targetRepository  repositories.TargetRepository
	sqlmapClient      clients.SqlmapClient
}

func NewScannerService(
	segmentRepository repositories.SegmentRepository,
	vulRepository repositories.VulRepository,
	targetRepository repositories.TargetRepository,
	sqlMapClient clients.SqlmapClient) *ScannerService {
	return &ScannerService{
		segmentRepository: segmentRepository,
		vulRepository:     vulRepository,
		targetRepository:  targetRepository,
		sqlmapClient:      sqlMapClient,
	}
}

// TODO restrict to scan one time for a segment
func (ss *ScannerService) ScanSegment(scanSegmentPack *packages.ScanSegmentPack, currentUser *models.User,
) (*models.Segment, enums.Error) {
	segment, err := ss.segmentRepository.FindByID(scanSegmentPack.SegmentID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrInvalidRequest
	}
	if err != nil {
		log.Printf("Error when get segment by ID: %s", scanSegmentPack.SegmentID)
		return nil, enums.ErrSystem
	}
	//if segment.ScanningStatus == enums.StatusTerminated {
	//	return nil, enums.NotAllowed
	//}
	if segment.UserID != currentUser.ID {
		return nil, enums.ErrUnauthorized
	}
	targets, err := ss.targetRepository.GetAllBySegmentID(segment.ID)
	if err != nil && err != enums.ErrEntityNotFound {
		return nil, enums.ErrSystem
	}
	if targets == nil || err == enums.ErrEntityNotFound {
		return nil, enums.ErrNoResources
	}
	go ss.ScanMultiTargets(targets)
	updateSegmentMap := make(map[string]interface{})
	updateSegmentMap["id"] = segment.ID
	updateSegmentMap["ScanningStatus"] = enums.StatusRunning
	updatedSegment, err := ss.segmentRepository.UpdateWithMap(updateSegmentMap)
	if err != nil {
		log.Printf("Can not update scanning status for segment with error: %s", err)
		return updatedSegment, enums.ErrSystem
	}
	return updatedSegment, nil
}

func (ss *ScannerService) ScanMultiTargets(targets []models.Target) {
	resultChan := make(chan *models.Vul)
	var wg sync.WaitGroup
	go func(segmentID string) {
		numVuls := 0
		for vul := range resultChan {
			err := ss.vulRepository.Create(vul)
			if err != nil {
				log.Printf("Can not save a vul with error: %s", err)
			}
			numVuls++
			log.Printf("Detect and save vul %s for segment %s", vul.ID, segmentID)
		}
		log.Printf("Save all vul of segment: %s", segmentID)
		updateSegmentMap := make(map[string]interface{})
		updateSegmentMap["id"] = segmentID
		updateSegmentMap["ScanningStatus"] = enums.StatusTerminated
		updateSegmentMap["VulNumber"] = numVuls
		_, err := ss.segmentRepository.UpdateWithMap(updateSegmentMap)
		if err != nil {
			log.Printf("Can not update scanning status segment %s after finish scanwiht error: %s",
				segmentID, err,
			)
		}
	}(targets[0].SegmentID)

	// TODO sqlmapapi is not strong enough to receive too much request
	for i := range targets {
		wg.Add(1)
		go func(target models.Target) {
			_ = ss.scanTarget(target, resultChan, &wg)
		}(targets[i])
	}
	wg.Wait()
	close(resultChan)
}

func (ss *ScannerService) scanTarget(target models.Target, resultChan chan *models.Vul, wg *sync.WaitGroup) error {
	defer wg.Done()
	var err error
	taskID, err := ss.sqlmapClient.NewTask()
	if err != nil {
		log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
		return enums.ErrSystem
	}
	log.Printf("Create new task: %s", taskID)
	switch target.Method {
	case http.MethodGet:
		err = ss.sqlmapClient.SetOptionGET(taskID)
	case http.MethodPost:
		mapParams := target.GetMapParams()
		err = ss.sqlmapClient.SetOptionForPOST(taskID, mapParams)
	}
	if err != nil {
		log.Printf("Error when scan [TargetID] %s [Error] %s", target.ID, err)
		return enums.ErrSystem
	}
	err = ss.sqlmapClient.StartScan(taskID, target.RawURL)
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
		//if status == enums.StatusRunning {
		//	time.Sleep(5 * time.Second)
		//}
		if status == enums.StatusTerminated {
			break
		}
		time.Sleep(5 * time.Second)
	}
	_, isVul, err := ss.sqlmapClient.GetData(taskID)
	if err != nil {
		log.Printf("Error when get result of task [TargetID: %s] [TaskID: %s] [Error: %s]", target.ID, taskID, err)
		return err
	}
	if isVul == enums.ResultExistVul {
		vul := models.NewVulWithTarget(target)
		log.Printf("Get new vul %v for target %s", vul, target.ID)
		resultChan <- vul
	}
	return nil
}
