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

func NewProjectRepository(baseRepository baseRepository) *ProjectRepository {
	return &ProjectRepository{baseRepository: baseRepository}
}

func (pr *ProjectRepository) FindProjectById(id string) (*models.Project, error) {
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
	err := pr.db.Create(project).Error
	if err != nil {
		return err
	}
	return nil
}

func (pr *ProjectRepository)
