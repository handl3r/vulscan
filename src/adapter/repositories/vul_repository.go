package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type VulRepository struct {
	baseRepository
}

func NewVulRepository(db *gorm.DB) *VulRepository {
	return &VulRepository{baseRepository: baseRepository{
		db: db,
	}}
}

func (vp *VulRepository) FindByID(id string) (*models.Vul, error) {
	vul := &models.Vul{}
	err := vp.db.Model(&models.Vul{}).Where("id = ?", id).Take(vul).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return vul, err
}

func (vp *VulRepository) GetBySegmentID(segmentID string) ([]models.Vul, error) {
	var vuls []models.Vul
	err := vp.db.Where("segment_id = ?", segmentID).Find(&vuls).Error
	if err == gorm.ErrRecordNotFound {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return vuls, nil
}

func (vp *VulRepository) Create(vul *models.Vul) error {
	vul.ID = uuid.New().String()
	err := vp.db.Model(&models.Vul{}).Create(vul).Error
	if err != nil {
		return err
	}
	return nil
}

func (vp *VulRepository) DeleteSegmentVuls(segmentID string) error {
	var vuls []models.Vul
	err := vp.db.Model(&models.Vul{}).Where("segment_id = ?", segmentID).Find(&vuls).Delete(&vuls).Error
	if err != nil {
		return err
	}
	return nil
}
