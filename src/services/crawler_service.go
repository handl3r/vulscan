package services

import (
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/ports"
)

type CrawlerService struct {
	baseService
	crawlerDriver ports.CrawlerPort
}

func NewCrawlerService(crawlerDriver ports.CrawlerPort) *CrawlerService {
	return &CrawlerService{
		baseService:   baseService{},
		crawlerDriver: crawlerDriver,
	}
}

func (c *CrawlerService) DiscoverURLs(domain string, typeLoadSite int) ([]models.Target, enums.Error) {
	return nil, nil
}
