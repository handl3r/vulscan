package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"log"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type SegmentRepository struct {
	baseRepository
}

func NewSegmentRepository(db *gorm.DB) *SegmentRepository {
	return &SegmentRepository{baseRepository: baseRepository{db: db}}
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

func (sp *SegmentRepository) GetByProjectID(projectID string) ([]models.Segment, error) {
	segments := make([]models.Segment, 0)
	err := sp.db.Order("created_at desc").Where("project_id = ?", projectID).Find(&segments).Error
	if err != nil {
		return nil, err
	}
	return segments, nil
}

func (sp *SegmentRepository) Create(segment *models.Segment) error {
	segment.ID = uuid.New().String()
	err := sp.db.Model(&models.Segment{}).Create(segment).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) UpdateWithMap(mapSegment map[string]interface{}) (*models.Segment, error) {
	err := sp.db.Model(&models.Segment{}).Where("id= ?", mapSegment["id"]).Updates(mapSegment).Error
	if err != nil {
		return nil, err
	}
	updatedSegment, err := sp.FindByID(mapSegment["id"].(string))
	if err != nil {
		log.Printf("Can not get segment by id with error: %s", err)
		return updatedSegment, err
	}
	return updatedSegment, nil
}

func (sp *SegmentRepository) UpdateIsCrawling(segmentID string, isCrawling bool) error {
	err := sp.db.Model(&models.Segment{}).Where("id=?", segmentID).Update("is_crawling", isCrawling).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) DeleteByID(id string) error {
	err := sp.db.Delete(&models.Segment{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (sp *SegmentRepository) DeleteProjectSegments(projectID string) error {
	var segments []models.Segment
	err := sp.db.Model(&models.Segment{}).Where("project_id = ?", projectID).Find(&segments).Delete(&segments).Error
	if err != nil {
		return err
	}
	return nil
}
