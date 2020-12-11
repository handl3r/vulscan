package crawler_driver

import (
	"vulscan/src/models"
)

// CollyCrawler is crawler for static page html
type CollyCrawler struct {
}

func (c CollyCrawler) CrawlURLs(domain string, maxDepth int) ([]models.Target, error) {
	//targets := make([]models.Target, 0)
	//collector := colly.NewCollector(
	//	colly.AllowedDomains(domain),
	//	colly.MaxDepth(maxDepth),
	//	)
	//collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
	//	link := element.Attr()
	//})

	return nil, nil

}
