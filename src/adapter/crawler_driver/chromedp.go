package crawler_driver

import (
	"vulscan/src/models"
)

// TODO make decorator for cover crawler to choose crawl by chromedp or colly
type ChromeDPCrawler struct {

}

func (c ChromeDPCrawler) CrawlURLs(domain string, maxDepth int) ([]models.Target, error) {
	panic("implement me")
}

