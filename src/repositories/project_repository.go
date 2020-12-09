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

func (pr *ProjectRepository) GetByUserID(userID string) ([]*models.Project, error) {
	user := &models.User{
		ID: userID,
	}
	projects := make([]*models.Project, 0)
	err := pr.db.Model(user).Association("UserID").Find(&projects)
	// recheck latter if Find return ErrRecordNotFound
	if err != nil {
		return nil, err
	}
	if pr.db.RowsAffected == 0 {
		return nil, enums.ErrEntityNotFound
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
	err := pr.db.Model(&models.Project{}).Delete(&models.Project{}, id).Error
	if err != nil {
		return err
	}
	return nil
}

// Update Find by
func (pr *ProjectRepository) Update(project *models.Project) (*models.Project, error) {
	err := pr.db.Model(&models.Project{}).Updates(project).Error
	if err != nil {
		return nil, err
	}
	existProject, err := pr.FindProjectByID(project.ID)
	if err != nil {
		return nil, err
	}
	return existProject, nil
}
