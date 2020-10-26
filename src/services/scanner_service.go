package services

import (
	"vulscan/src/models"
	"vulscan/src/repositories"
)

type ScannerService struct {
	segmentRepository repositories.SegmentRepository
}

func (ss *ScannerService) Scan([]models.Target) ([]models.Vul, error) {
	return nil, nil
}
