package services

import (
	"log"
	"net/http"
	"time"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/helpers"
	"vulscan/src/models"
	"vulscan/src/packages"
)

type RegistrationService struct {
	userRepository *repositories.UserRepository
}

func NewRegistrationService(userRepository *repositories.UserRepository) *RegistrationService {
	return &RegistrationService{
		userRepository: userRepository,
	}
}

func (r *RegistrationService) Signup(pack packages.RegistrationPack) (*models.User, enums.Error) {
	_, err := r.validateRegistration(pack)
	if err != nil {
		return nil, err
	}
	hashedPassword, newErr := helpers.HashPassword(pack.Password)
	if newErr != nil {
		return nil, enums.ErrSystem
	}
	user := &models.User{
		Email:     pack.Email,
		Password:  hashedPassword,
		CreatedAt: time.Now(),
	}
	newErr = r.userRepository.Create(user)
	if newErr != nil {
		return nil, enums.ErrSystem
	}
	return user, nil
}

func (r *RegistrationService) validateRegistration(pack packages.RegistrationPack) (bool, enums.Error) {
	if pack.Password != pack.PasswordConfirmation {
		return false, enums.NewHttpCustomError(
			http.StatusConflict,
			"un-match_password",
			"Un-match password with password confirmation",
		)
	}
	if !enums.EmailRegex.MatchString(pack.Email) {
		return false, enums.NewHttpCustomError(
			http.StatusBadRequest,
			"invalid_email",
			"Invalid email",
		)
	}
	_, err := r.userRepository.FindByEmail(pack.Email)
	if err == nil {
		return false, enums.NewHttpCustomError(
			http.StatusConflict,
			"email_unique",
			"This email is already registered",
		)
	}
	if err != enums.ErrEntityNotFound {
		log.Printf("Can not validate registration package with error: %s", err)
		return false, enums.ErrSystem
	}
	return true, nil
}
