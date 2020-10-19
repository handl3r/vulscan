package repositories

import (
	"errors"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type VulInfoRepository struct {
	baseRepository
}

func NewVulInfoRepository(baseRepository baseRepository) *VulInfoRepository {
	return &VulInfoRepository{baseRepository: baseRepository}
}

func (vir *VulInfoRepository) GetByID(id string) (*models.VulInfo, error) {
	vulInfo := &models.VulInfo{}
	err := vir.db.Model(&models.VulInfo{}).Where("id = ?", id).Take(vulInfo).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return vulInfo, nil
}

func (vir *VulInfoRepository) Create(vulInfo *models.VulInfo) error {
	err := vir.db.Model(&models.VulInfo{}).Create(vulInfo).Error
	if err != nil {
		return err
	}
	return nil
}

func (vir *VulInfoRepository) DeleteVulVulInfos(vulID string) error {
	vulInfos := make([]models.VulInfo, 0)

	err := vir.db.Model(&models.VulInfo{}).Where("vul_id = ?", vulID).Find(vulInfos).Delete(vulInfos).Error
	if err != nil {
		return err
	}
	return nil
}
