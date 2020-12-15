package bootstrap

import (
	"vulscan/api/http/context"
	"vulscan/configs"
	"vulscan/src/adapter/clients"
	"vulscan/src/adapter/crawler"
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

	collyCrawler := crawler.NewCollyCrawler(conf.MaximumTargetCrawler)
	chromeDPCrawler := crawler.NewChromeDPCrawler()
	crawlerDriver := crawler.NewCrawlerDecorator(chromeDPCrawler, collyCrawler)
	crawlerService := services.NewCrawlerService(crawlerDriver, targetRepository, segmentRepository)

	sqlMapClient := clients.NewSqlmapClient("", conf.SQLMapServerHost, conf.SQLMapServerPort)

	projectService := services.NewProjectService(projectRepository, segmentRepository, targetRepository, crawlerService)
	userService := services.NewUserService(*userRepository, *projectRepository)
	segmentService := services.NewSegmentService(*segmentRepository, *targetRepository, *vulRepository)
	authService := services.NewAuthenticationService(userRepository, conf.AccessTokenExp, conf.AuthSecretKet)
	registrationService := services.NewRegistrationService(userRepository)
	scannerService := services.NewScannerService(*segmentRepository, *vulRepository, *targetRepository, *sqlMapClient)
	appContext := context.ApplicationContext{
		ProjectService:      projectService,
		UserService:         userService,
		SegmentService:      segmentService,
		AuthService:         authService,
		RegistrationService: registrationService,
		ScannerService:      scannerService,
		DBConnection:        dbConnection,
	}
	return &appContext
}
