package context

import (
	"gorm.io/gorm"
	"vulscan/src/services"
)

type ApplicationContext struct {
	ProjectService      *services.ProjectService
	UserService         *services.UserService
	SegmentService      *services.SegmentService
	AuthService         *services.AuthenticationService
	RegistrationService *services.RegistrationService
	DBConnection        *gorm.DB
}
