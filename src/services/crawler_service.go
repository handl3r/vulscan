package services

import (
	"log"
	"net/http"
	"net/url"
	"vulscan/src/adapter/repositories"
	"vulscan/src/common"
	"vulscan/src/enums"
	"vulscan/src/models"
	"vulscan/src/packages"
)

type CrawlerService struct {
	baseService
	crawlerDriver     common.CrawlerPort
	targetRepository  *repositories.TargetRepository
	segmentRepository *repositories.SegmentRepository
}

func NewCrawlerService(crawlerDriver common.CrawlerPort, targetRepository *repositories.TargetRepository,
	segmentRepository *repositories.SegmentRepository) *CrawlerService {
	return &CrawlerService{
		baseService:       baseService{},
		crawlerDriver:     crawlerDriver,
		targetRepository:  targetRepository,
		segmentRepository: segmentRepository,
	}
}

func (c *CrawlerService) DiscoverURLs(domain *url.URL, typeLoadSite string, maxDepth int) ([]models.Target, enums.Error) {
	targets, err := c.crawlerDriver.CrawlURLs(domain, typeLoadSite, maxDepth)
	if err != nil {
		return nil, enums.ErrSystem
	}
	return targets, nil
}

func (c *CrawlerService) DiscoverURLsAndSave(discoverPack *packages.DiscoverProjectPack, segmentID string, project *models.Project) ([]models.Target, enums.Error) {
	valid, err := c.validateDiscoverPack(discoverPack)
	if !valid {
		return nil, err
	}
	parsedURL, newErr := url.Parse(project.Domain)
	if newErr != nil {
		log.Printf("Can not parse domain %s of project %s to prepare for discover urls with error: %s",
			project.Domain, project.ID, err)
		return nil, enums.ErrSystem
	}
	targets, err := c.DiscoverURLs(parsedURL, discoverPack.IsLoadByJS, discoverPack.DepthLevel)
	if err != nil {
		log.Printf("Can not crawl project %s, domain %s with error: %s", project.ID, project.Domain, err)
		return nil, enums.ErrSystem
	}
	for i := range targets {
		targets[i].SegmentID = segmentID
	}
	newError := c.targetRepository.SaveTargets(targets)
	if newError != nil {
		log.Printf("Can not save targers of segment %s in project %s with error: %s", segmentID, project.ID, newError)
		return nil, enums.ErrSystem
	}
	updateMap := make(map[string]interface{})
	updateMap["id"] = segmentID
	updateMap["IsCrawling"] = false
	updateMap["TargetNumber"] = len(targets)
	_, newError = c.segmentRepository.UpdateWithMap(updateMap)
	if newError != nil {
		log.Printf("Can not update crawling status for segment %s", segmentID)
	}
	return targets, nil
}

func (c *CrawlerService) validateDiscoverPack(discoverPack *packages.DiscoverProjectPack) (bool, enums.Error) {
	if len(discoverPack.ProjectID) == 0 {
		return false, enums.NewHttpCustomError(
			http.StatusBadRequest,
			"invalid_project_id",
			"Invalid project id",
		)
	}
	if ok := enums.DiscoverDepthMap[discoverPack.DepthLevel]; !ok {
		return false, enums.NewHttpCustomError(
			http.StatusBadRequest,
			"invalid_depth_level",
			"Invalid depth level",
		)
	}
	if discoverPack.IsLoadByJS != enums.TypeStaticSite && discoverPack.IsLoadByJS != enums.TypeDynamicSite {
		return false, enums.NewHttpCustomError(
			http.StatusBadRequest,
			"invalid_is_load_by_js",
			"Invalid type load site",
		)
	}
	return true, nil
}
