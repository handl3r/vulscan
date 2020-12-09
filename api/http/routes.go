package http

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/controllers"
)

type ControllerManager struct {
	UserController    *controllers.UserController
	ProjectController *controllers.ProjectController
}

func NewRouter(controllerManager *ControllerManager) *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1/user")
	v1.GET("/projects", controllerManager.UserController.GetProjectsByUserID)

	v2 := router.Group("/api/v1/projects")
	v2.GET("/:id", controllerManager.ProjectController.Get)
	v2.POST("", controllerManager.ProjectController.Create)
	v2.PATCH("/:id", controllerManager.ProjectController.Update)
	v2.DELETE("/:id", controllerManager.ProjectController.Delete)

	return router
}
