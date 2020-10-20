package services

import (
	"errors"
	"gorm.io/gorm"
	"log"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/repositories"
)

type ProjectService struct {
	projectRepository *repositories.ProjectRepository
	segmentRepository *repositories.SegmentRepository
}

// GetAll get all project basic information of a user
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

// GetByID get project information and all segments belongs to it
func (ps *ProjectService) GetByID(projectID string) (*models.Project, enums.Error) {
	project, err := ps.projectRepository.FindProjectByID(projectID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	return project, nil
}

func (ps *ProjectService) Delete(projectID string) enums.Error {
	project, err := ps.projectRepository.FindProjectByID(projectID)
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return enums.ErrInvalidRequest
	}
	if err != nil {
		return enums.ErrSystem
	}
	err = ps.segmentRepository.DeleteProjectSegments(project.ID)
	if err != nil {
		log.Printf("[E] Can not delete segment in project: %s", err)
	}
	return nil
}
