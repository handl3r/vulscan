package controllers

import "vulscan/src/services"

type ApplicationContext struct {
	ProjectService *services.ProjectService
	UserService    *services.UserService
	SegmentService *services.SegmentService
}
