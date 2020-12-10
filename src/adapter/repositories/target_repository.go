package repositories

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type TargetRepository struct {
	baseRepository
}

func NewTargetRepository(baseRepository baseRepository) *TargetRepository {
	return &TargetRepository{
		baseRepository,
	}
}

func (t *TargetRepository) GetAllBySegmentID(segmentID string) ([]models.Target, error) {
	targets := make([]models.Target, 0)
	err := t.db.Model(&models.Target{}).Where("segment_id = ?", segmentID).Find(&targets).Error
	if err == gorm.ErrRecordNotFound {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return targets, nil
}

func (t *TargetRepository) SaveTargets(targets []models.Target) error {
	for i := range targets {
		targets[i].ID = uuid.New().String()
	}
	err := t.db.Create(&targets).Error
	if err != nil {
		return err
	}
	return nil
}
