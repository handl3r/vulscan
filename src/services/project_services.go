package services

import (
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/repositories"
)

type ProjectService struct {
	projectRepository *repositories.ProjectRepository
}

// GetAll get all project of a user
func (ps *ProjectService) GetAll(userID string) ([]*models.Project, enums.Error) {
	projects, err := ps.projectRepository.GetByUserID(userID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrNoResources
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	return projects, nil
}

//func (ps *ProjectService) GetByID(projectID string) (*models.Project, enums.Error) {
//	project, err := ps.projectRepository.FindProjectByID(projectID)
//	if err == enums.ErrEntityNotFound {
//		return nil, enums.ErrResourceNotFound
//	}
//	if err != nil {
//		return nil, enums.ErrSystem
//	}
//
//
//}

