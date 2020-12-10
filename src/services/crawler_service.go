package services

import (
	"vulscan/src/enums"
	"vulscan/src/models"
)

type CrawlerService struct {
	baseService
}

func NewCrawlerService() *CrawlerService {
	return &CrawlerService{
		baseService{},
	}
}

func (c *CrawlerService) DiscoverURLs(domain string) ([]models.Target, enums.Error) {
	return nil, nil
}
