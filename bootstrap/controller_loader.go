package bootstrap

import (
	"vulscan/api/http"
	"vulscan/api/http/context"
	"vulscan/api/http/controllers"
)

func LoadControllerManager(appContext *context.ApplicationContext) *http.ControllerManager {
	userController := controllers.NewUserController(appContext)
	projectController := controllers.NewProjectController(appContext)
	segmentController := controllers.NewSegmentController(appContext)
	authController := controllers.NewAuthenticationController(appContext)
	scannerController := controllers.NewScannerController(appContext)

	return &http.ControllerManager{
		UserController:    userController,
		ProjectController: projectController,
		SegmentController: segmentController,
		AuthController:    authController,
		ScannerController: scannerController,
	}
}
