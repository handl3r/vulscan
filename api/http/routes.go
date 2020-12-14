package http

import (
	"github.com/gin-gonic/gin"
	"vulscan/api/http/context"
	"vulscan/api/http/controllers"
	"vulscan/api/http/middlewares"
)

type ControllerManager struct {
	UserController    *controllers.UserController
	ProjectController *controllers.ProjectController
	SegmentController *controllers.SegmentController
	AuthController    *controllers.AuthenticationController
}

func NewRouter(context *context.ApplicationContext, controllerManager *ControllerManager) *gin.Engine {
	router := gin.Default()

	router.Group("/api/v1/user").
		GET("/projects", middlewares.RequireAccessToken(context), controllerManager.UserController.GetProjectsByUserID)

	router.Group("/api/v1/projects").
		GET("/:id", middlewares.RequireAccessToken(context), controllerManager.ProjectController.Get).
		POST("", middlewares.RequireAccessToken(context), controllerManager.ProjectController.Create).
		PATCH("/:id", middlewares.RequireAccessToken(context), controllerManager.ProjectController.Update).
		DELETE("/:id", middlewares.RequireAccessToken(context), controllerManager.ProjectController.Delete)

	router.Group("/api/v1/segments").
		GET("/:id", middlewares.RequireAccessToken(context), controllerManager.SegmentController.Get).
		DELETE("/:id", middlewares.RequireAccessToken(context), controllerManager.SegmentController.Delete)

	router.
		POST("/api/v1/discover", middlewares.RequireAccessToken(context), controllerManager.ProjectController.Discover).
		POST("/api/v1/scan", middlewares.RequireAccessToken(context), controllerManager.SegmentController.Get)
	router.POST("/api/v1/signup", controllerManager.AuthController.Register)
	router.POST("/api/v1/login", controllerManager.AuthController.Login)
	return router
}
