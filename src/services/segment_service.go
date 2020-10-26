package services

import (
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/repositories"
)

type SegmentService struct {
	segmentRepository repositories.SegmentRepository
	targetRepository  repositories.TargetRepository
	vulRepository     repositories.VulRepository
}

func NewSegmentService(segmentRepository repositories.SegmentRepository, targetRepository repositories.TargetRepository, vulRepository repositories.VulRepository) *SegmentService {
	return &SegmentService{segmentRepository: segmentRepository, targetRepository: targetRepository, vulRepository: vulRepository}
}

func (ss *SegmentService) GetByID(id string) (*models.Segment, error) {
	segment, err := ss.segmentRepository.FindByID(id)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, err
	}
	targets, err := ss.targetRepository.GetAllBySegmentID(segment.ID)
	if err != nil {
		return segment, err
	}
	segment.Targets = targets
	vuls, err := ss.vulRepository.GetBySegmentID(segment.ID)
	if err != nil {
		return segment, err
	}
	segment.Vuls = vuls
	return segment, nil
}

//func (ss *SegmentService) Create(segmentPack )
