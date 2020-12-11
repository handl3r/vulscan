package crawler_driver

import (
	"net/url"
	"vulscan/src/models"
)

// TODO make decorator for cover crawler to choose crawl by chromedp or colly
type ChromeDPCrawler struct {

}

func (c ChromeDPCrawler) CrawlURLs(domain *url.URL, maxDepth int) ([]models.Target, error) {
	panic("implement me")
}

