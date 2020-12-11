package services

import (
	"log"
	"vulscan/src/adapter/repositories"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/ports"
)

type CrawlerService struct {
	baseService
	crawlerDriver     ports.CrawlerPort
	targetRepository  *repositories.TargetRepository
	segmentRepository *repositories.SegmentRepository
}

func NewCrawlerService(crawlerDriver ports.CrawlerPort) *CrawlerService {
	return &CrawlerService{
		baseService:   baseService{},
		crawlerDriver: crawlerDriver,
	}
}

func (c *CrawlerService) DiscoverURLs(domain string, typeLoadSite string) ([]models.Target, enums.Error) {
	targets, err := c.crawlerDriver.CrawlURLs(domain, typeLoadSite, enums.DefaultMaxDepth)
	if err != nil {
		return nil, enums.ErrSystem
	}
	return targets, nil
}

func (c *CrawlerService) DiscoverURLsAndSave(domain, typeLoadSite, segmentID, projectID string) ([]models.Target, enums.Error) {
	targets, err := c.DiscoverURLs(domain, typeLoadSite)
	if err != nil {
		log.Printf("Can not crawl project %s, domain %s with error: %s", projectID, domain, err)
		return nil, enums.ErrSystem
	}
	for i := range targets {
		targets[i].SegmentID = segmentID
	}
	newError := c.targetRepository.SaveTargets(targets)
	if newError != nil {
		log.Printf("Can not save targers of segment %s in project %s with error: %s", segmentID, projectID, newError)
		return nil, enums.ErrSystem
	}
	newSegment := &models.Segment{
		ID:         segmentID,
		IsCrawling: false,
	}
	newError = c.segmentRepository.Update(newSegment)
	if newError != nil {
		log.Printf("Can not update crawling status for segment %s", segmentID)
	}
	return targets, nil
}
