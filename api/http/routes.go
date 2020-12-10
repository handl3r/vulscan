package http

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/controllers"
)

type ControllerManager struct {
	UserController    *controllers.UserController
	ProjectController *controllers.ProjectController
	SegmentController *controllers.SegmentController
}

func NewRouter(controllerManager *ControllerManager) *gin.Engine {
	router := gin.Default()
	router.Group("/api/v1/user").
		GET("/projects", controllerManager.UserController.GetProjectsByUserID)

	router.Group("/api/v1/projects").
		GET("/:id", controllerManager.ProjectController.Get).
		POST("", controllerManager.ProjectController.Create).
		PATCH("/:id", controllerManager.ProjectController.Update).
		DELETE("/:id", controllerManager.ProjectController.Delete)

	router.Group("/api/v1/segments").
		GET("/:id", controllerManager.SegmentController.Get).
		DELETE("/:id", controllerManager.SegmentController.Delete)

	return router
}
