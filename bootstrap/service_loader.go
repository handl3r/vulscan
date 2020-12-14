package bootstrap

import (
	"vulscan/api/http/context"
	"vulscan/configs"
	"vulscan/src/adapter/crawler_driver"
	"vulscan/src/adapter/repositories"
	"vulscan/src/services"
)

func LoadServices(conf *configs.Config) *context.ApplicationContext {
	dbConnection := initDBConnection(conf)

	projectRepository := repositories.NewProjectRepository(dbConnection)
	segmentRepository := repositories.NewSegmentRepository(dbConnection)
	targetRepository := repositories.NewTargetRepository(dbConnection)
	userRepository := repositories.NewUserRepository(dbConnection)
	vulRepository := repositories.NewVulRepository(dbConnection)

	collyCrawler := crawler_driver.NewCollyCrawler(conf.MaximumTargetCrawler)
	chromeDPCrawler := crawler_driver.NewChromeDPCrawler()
	crawlerDriver := crawler_driver.NewCrawlerDecorator(chromeDPCrawler, collyCrawler)
	crawlerService := services.NewCrawlerService(crawlerDriver, targetRepository, segmentRepository)

	projectService := services.NewProjectService(projectRepository, segmentRepository, targetRepository, crawlerService)
	userService := services.NewUserService(*userRepository, *projectRepository)
	segmentService := services.NewSegmentService(*segmentRepository, *targetRepository, *vulRepository)
	authService := services.NewAuthenticationService(userRepository, conf.AccessTokenExp, conf.AuthSecretKet)
	registrationService := services.NewRegistrationService(userRepository)
	appContext := context.ApplicationContext{
		ProjectService:      projectService,
		UserService:         userService,
		SegmentService:      segmentService,
		AuthService:         authService,
		RegistrationService: registrationService,
		DBConnection:        dbConnection,
	}
	return &appContext
}
