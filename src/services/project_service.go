package services

import (
	"log"
	url2 "net/url"
	"time"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/packages"
)

type ProjectService struct {
	projectRepository *repositories.ProjectRepository
	segmentRepository *repositories.SegmentRepository
	targetRepository  *repositories.TargetRepository
	crawlerService    *CrawlerService
}

func NewProjectService(projectRepository *repositories.ProjectRepository, segmentRepository *repositories.SegmentRepository,
	targetRepository *repositories.TargetRepository, crawlerService *CrawlerService,
) *ProjectService {
	return &ProjectService{
		projectRepository: projectRepository,
		segmentRepository: segmentRepository,
		targetRepository:  targetRepository,
		crawlerService:    crawlerService,
	}
}

// Create create a project
func (ps *ProjectService) Create(pack *packages.CreateProjectPack, currentUser *models.User,
) (*models.Project, enums.Error) {
	if !packages.ValidateCreateProjectPack(pack) {
		return nil, enums.ErrInvalidRequest
	}
	projectModel := &models.Project{
		Name:   pack.Name,
		Domain: pack.Domain,
		UserID: currentUser.ID,
	}
	err := ps.projectRepository.Create(projectModel)
	log.Printf("Can not create project for currentUser %s with error: %s", currentUser.ID, err)
	if err != nil {
		return nil, enums.ErrSystem
	}
	return projectModel, nil
}

func (ps *ProjectService) Update(pack *packages.UpdateProjectPack, currentUser *models.User) (*models.Project, enums.Error) {
	projectModel := &models.Project{
		ID:   pack.ID,
		Name: pack.Name,
	}
	existedProject, err := ps.projectRepository.FindProjectByID(pack.ID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	if existedProject.UserID != currentUser.ID {
		return nil, enums.ErrResourceNotFound
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
func (ps *ProjectService) GetByID(projectID string, currentUser *models.User) (*models.Project, enums.Error) {
	project, err := ps.projectRepository.FindProjectByID(projectID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	if project.UserID != currentUser.ID {
		return nil, enums.ErrResourceNotFound
	}
	segments, err := ps.segmentRepository.GetByProjectID(projectID)
	if err != nil {
		return project, enums.ErrSystem
	}
	project.Segments = segments
	return project, nil
}

// TODO change to event bus when delete to finish stuff jobs in background
func (ps *ProjectService) DeleteByID(projectID string, currentUser *models.User) enums.Error {
	project, err := ps.projectRepository.FindProjectByID(projectID)
	if err == enums.ErrEntityNotFound {
		return enums.ErrResourceNotFound
	}
	if err != nil {
		return enums.ErrSystem
	}
	if project.UserID != currentUser.ID {
		return enums.ErrResourceNotFound
	}
	go func() {
		err = ps.segmentRepository.DeleteProjectSegments(project.ID)
		if err != nil {
			log.Printf("[E] Can not delete segment in project %s with error: %s", project.ID, err)
		}
	}()
	return nil
}

// TODO change to response 200 after create segment and crawl in background
func (ps *ProjectService) Crawl(discoverProjectPack *packages.DiscoverProjectPack, currentUser *models.User) (*models.Segment, enums.Error) {
	project, err := ps.projectRepository.FindProjectByID(discoverProjectPack.ProjectID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrResourceNotFound
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	if project.UserID != currentUser.ID {
		return nil, enums.ErrResourceNotFound
	}
	now := time.Now()
	segment := &models.Segment{
		IsScanned:  false,
		IsCrawling: true,
		CreatedAt:  now,
		ProjectID:  project.ID,
		UserID:     currentUser.ID,
	}
	err = ps.segmentRepository.Create(segment)
	if err != nil {
		return nil, enums.ErrSystem
	}
	go func() {
		url, err := url2.Parse(project.Domain)
		if err != nil {
			log.Printf("Can not parse domain %s of project %s to prepare for discover urls with error: %s",
				project.Domain, project.ID, err)
			return
		}
		_, _ = ps.crawlerService.DiscoverURLsAndSave(url, discoverProjectPack.IsLoadByJS, segment.ID, project.ID)
	}()
	return segment, nil
}
