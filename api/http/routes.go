package http

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/controllers"
	"vulscan/api/http/middlewares"
)

type ControllerManager struct {
	UserController    *controllers.UserController
	ProjectController *controllers.ProjectController
	SegmentController *controllers.SegmentController
	AuthController    *controllers.AuthenticationController
}

func NewRouter(context *ApplicationContext, controllerManager *ControllerManager) *gin.Engine {
	router := gin.Default()

	router.POST("/api/v1/signup", controllerManager.AuthController.Register)
	router.POST("/api/v1/login", controllerManager.AuthController.Login)

	router.Group("/api/v1/user").
		Use(middlewares.RequireAccessToken(context)).
		GET("/projects", controllerManager.UserController.GetProjectsByUserID)

	router.Group("/api/v1/projects").
		Use(middlewares.RequireAccessToken(context)).
		GET("/:id", controllerManager.ProjectController.Get).
		POST("", controllerManager.ProjectController.Create).
		PATCH("/:id", controllerManager.ProjectController.Update).
		DELETE("/:id", controllerManager.ProjectController.Delete)

	router.Group("/api/v1/segments").
		Use(middlewares.RequireAccessToken(context)).
		GET("/:id", controllerManager.SegmentController.Get).
		DELETE("/:id", controllerManager.SegmentController.Delete)

	router.Use(middlewares.RequireAccessToken(context)).
		POST("/api/v1/discover", controllerManager.ProjectController.Discover)
	return router
}
