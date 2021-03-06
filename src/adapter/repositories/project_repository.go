package repositories

import (
	"errors"
	"github.com/google/uuid"
	"gorm.io/gorm"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type ProjectRepository struct {
	baseRepository
}

func NewProjectRepository(db *gorm.DB) *ProjectRepository {
	return &ProjectRepository{baseRepository: baseRepository{db: db}}
}

func (pr *ProjectRepository) GetByUserID(userID string) ([]*models.Project, error) {
	var projects []*models.Project
	err := pr.db.Where("user_id = ?", userID).Find(&projects).Error
	if err == gorm.ErrRecordNotFound {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return projects, nil
}

func (pr *ProjectRepository) FindProjectByID(id string) (*models.Project, error) {
	project := &models.Project{}
	err := pr.db.Model(&models.Project{}).Where("id = ?", id).Take(project).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, enums.ErrEntityNotFound
	}
	if err != nil {
		return nil, err
	}
	return project, nil
}

func (pr *ProjectRepository) Create(project *models.Project) error {
	project.ID = uuid.New().String()
	err := pr.db.Model(&models.Project{}).Create(project).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProjectRepository) DeleteByID(id string) error {
	err := pr.db.Delete(&models.Project{ID: id}).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProjectRepository) Update(project *models.Project) (*models.Project, error) {
	err := pr.db.Model(&models.Project{}).Where("id= ?", project.ID).Updates(project).Error
	if err != nil {
		return nil, err
	}
	updatedProject, err := pr.FindProjectByID(project.ID)
	if err != nil {
		return nil, err
	}
	return updatedProject, nil
}
