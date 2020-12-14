package services

import (
	"log"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type SegmentService struct {
	segmentRepository repositories.SegmentRepository
	targetRepository  repositories.TargetRepository
	vulRepository     repositories.VulRepository
}

func NewSegmentService(segmentRepository repositories.SegmentRepository, targetRepository repositories.TargetRepository,
	vulRepository repositories.VulRepository) *SegmentService {
	return &SegmentService{
		segmentRepository: segmentRepository,
		targetRepository:  targetRepository,
		vulRepository:     vulRepository,
	}
}

// GetByID get segment with all information by ID
func (s *SegmentService) GetByID(id string, currentUser *models.User) (*models.Segment, enums.Error) {
	segment, err := s.segmentRepository.FindByID(id)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	if segment.UserID != currentUser.ID {
		return nil, enums.ErrResourceNotFound
	}
	targets, err := s.targetRepository.GetAllBySegmentID(segment.ID)
	if err != nil && err != enums.ErrEntityNotFound {
		return segment, enums.ErrSystem
	}
	segment.Targets = targets
	vuls, err := s.vulRepository.GetBySegmentID(segment.ID)
	if err != nil && err != enums.ErrEntityNotFound {
		return segment, enums.ErrSystem
	}
	segment.Vuls = vuls
	return segment, nil
}

// DeleteByID delete a segment by ID
func (s *SegmentService) DeleteByID(id string, currentUser *models.User) enums.Error {
	segment, err := s.segmentRepository.FindByID(id)
	if err == enums.ErrEntityNotFound {
		return enums.ErrResourceNotFound
	}
	if err != nil {
		return enums.ErrSystem
	}
	if segment.UserID != currentUser.ID {
		return enums.ErrResourceNotFound
	}
	err = s.segmentRepository.DeleteByID(id)
	if err != nil {
		return enums.ErrSystem
	}
	go func() {
		err := s.vulRepository.DeleteSegmentVuls(id)
		if err != nil {
			log.Printf("[E] Can not delete vuls in segment %s with error: %s", segment.ID, err)
		}
	}()
	return nil
}
