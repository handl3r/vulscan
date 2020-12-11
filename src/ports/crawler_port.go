package ports

import "vulscan/src/models"

type CrawlerPort interface {
	CrawlURLs(domain string, typeLoadSite string, maxDepth int) ([]models.Target, error)
}
