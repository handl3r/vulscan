package ports

import (
	"net/url"
	"vulscan/src/models"
)

type CrawlerPort interface {
	CrawlURLs(domain *url.URL, typeLoadSite string, maxDepth int) ([]models.Target, error)
}
