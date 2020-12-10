package services

import (
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type UserService struct {
	userRepository    repositories.UserRepository
	projectRepository repositories.ProjectRepository
}

func NewUserService(userRepository repositories.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}

func (u *UserService) GetProjectByUser(currentUser *models.User) ([]*models.Project, enums.Error) {
	projects, err := u.projectRepository.GetByUserID(currentUser.ID)
	if err == enums.ErrEntityNotFound {
		return nil, enums.ErrNoResources
	}
	if err != nil {
		return nil, enums.ErrSystem
	}
	return projects, nil
}
