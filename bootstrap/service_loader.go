package bootstrap

import (
	"vulscan/api/http"
	"vulscan/configs"
	"vulscan/src/adapter/crawler_driver"
	"vulscan/src/services"
)

func LoadServices(conf *configs.Config) http.ApplicationContext {
	dbConnection := initDBConnection(conf)
	appContext := &http.ApplicationContext{
		ProjectService:
	}
	collyCrawler := crawler_driver.NewCollyCrawler()

	crawlerService := services.NewCrawlerService()
	projectService := services.NewProjectService()
}
