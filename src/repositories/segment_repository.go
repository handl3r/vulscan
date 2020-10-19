package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type SegmentRepository struct {
	baseRepository
}

func NewSegmentRepository(baseRepository baseRepository) *SegmentRepository {
	return &SegmentRepository{baseRepository: baseRepository}
}

func (sp *SegmentRepository) FindByID(id string) (*models.Segment, error) {
	segment := &models.Segment{}
	err := sp.db.Model(&models.Segment{}).Where("id = ?", id).Take(segment).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return segment, nil
}

func (sp *SegmentRepository) Create(segment *models.Segment) error {
	segment.ID = uuid.New().String()
	err := sp.db.Model(&models.Segment{}).Create(segment).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) Update(segment *models.Segment) error {
	err := sp.db.Model(&models.Segment{}).Save(segment).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) DeleteByID(id string) error {
	err := sp.db.Model(&models.Segment{}).Delete(&models.Segment{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) DeleteProjectSegments(projectID string) error {
	segments := make([]models.Segment, 0)
	err := sp.db.Model(&models.Segment{}).Where("project_id = ?", projectID).Find(segments).Delete(segments).Error
	if err != nil {
		return err
	}
	return nil
}
