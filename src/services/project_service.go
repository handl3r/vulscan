package services

import (
	"log"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/packages"
	"vulscan/src/repositories"
	"vulscan/src/validation"
)

type ProjectService struct {
	projectRepository *repositories.ProjectRepository
	segmentRepository *repositories.SegmentRepository
}

func NewProjectService(projectRepository *repositories.ProjectRepository, segmentRepository *repositories.SegmentRepository,
) *ProjectService {
	return &ProjectService{projectRepository: projectRepository, segmentRepository: segmentRepository}
}

// Create create a project
func (ps *ProjectService) Create(pack *packages.CreateProjectPack, user *models.User,
) (*models.Project, enums.Error) {
	if !validation.ValidateCreateProjectPack(pack) {
		return nil, enums.ErrInvalidRequest
	}
	projectModel := &models.Project{
		Name:   pack.Name,
		Domain: pack.Domain,
		User:   user,
	}
	err := ps.projectRepository.Create(projectModel)
	log.Printf("Can not create project for user %s with error: %s", user.ID, err)
	if err != nil {
		return nil, enums.ErrSystem
	}
	return projectModel, nil
}

func (ps *ProjectService) Update(pack *packages.UpdateProjectPack) (*models.Project, enums.Error) {
	projectModel := &models.Project{
		ID:   pack.ID,
		Name: pack.Name,
	}
	updatedProject, err := ps.projectRepository.Update(projectModel)
	if err != nil {
		log.Printf("Can not update project by ID %s", projectModel.ID)
		return nil, enums.ErrSystem
	}
	return updatedProject, nil
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
	segments, err := ps.segmentRepository.GetByProjectID(projectID)
	if err != nil {
		return project, enums.ErrSystem
	}
	project.Segments = segments
	return project, nil
}

// TODO change to event bus when delete to finish stuff jobs in background
func (ps *ProjectService) Delete(projectID string) enums.Error {
	project, err := ps.projectRepository.FindProjectByID(projectID)
	if err == enums.ErrEntityNotFound {
		return enums.ErrResourceNotFound
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
