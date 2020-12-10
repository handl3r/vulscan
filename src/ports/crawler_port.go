package ports

import "vulscan/src/models"

type CrawlerPort interface {
	CrawlURLs(domain string, typeLoadSite int, maxDepth int) ([]models.Target, error)
}
