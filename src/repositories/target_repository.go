package repositories

import "vulscan/src/models"

type TargetRepository struct {
	baseRepository
}

func NewTargetRepository(baseRepository baseRepository) *TargetRepository {
	return &TargetRepository{
		baseRepository,
	}
}

func (tr *TargetRepository) GetAllBySegmentID(segmentID string) ([]models.Target, error) {
	targets := make([]models.Target, 0)
	err := tr.db.Model(&models.Target{}).Where("segment_id = ?", segmentID).Find(&targets).Error
	if err != nil {
		return nil, err
	}
	return targets, nil
}
