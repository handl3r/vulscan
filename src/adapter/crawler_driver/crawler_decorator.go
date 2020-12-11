package crawler_driver

import (
	"net/url"
	"vulscan/src/enums"
	"vulscan/src/models"
)

type CrawlerDecorator struct {
	chromeDPCrawler *ChromeDPCrawler
	collyCrawler    *CollyCrawler
}

func NewCrawlerDecorator(chromeDPCrawler *ChromeDPCrawler, collyCrawler *CollyCrawler) *CrawlerDecorator {
	return &CrawlerDecorator{
		chromeDPCrawler: chromeDPCrawler,
		collyCrawler:    collyCrawler,
	}
}

func (c *CrawlerDecorator) CrawlURLs(domain *url.URL, typeLoadSite string, maxDepth int) ([]models.Target, error) {
	targets := make([]models.Target, 0)
	var err error
	switch typeLoadSite {
	case enums.TypeStaticSite:
		targets, err = c.collyCrawler.CrawlURLs(domain, maxDepth)
	case enums.TypeDynamicSite:
		targets, err = c.chromeDPCrawler.CrawlURLs(domain, maxDepth)
	}
	if err != nil {
		return nil, err
	}
	return targets, nil
}
