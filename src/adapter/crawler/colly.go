package crawler

import (
	"github.com/gocolly/colly"
	"log"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"vulscan/src/models"
)

// CollyCrawler is crawler for static page html
type CollyCrawler struct {
	baseCrawler
}

func NewCollyCrawler(maximumTarget int) *CollyCrawler {
	return &CollyCrawler{
		baseCrawler{
			maximumTarget: maximumTarget,
		},
	}
}

// TODO trigger maximum target
func (c *CollyCrawler) CrawlURLs(domain *url.URL, maxDepth int) ([]models.Target, error) {
	targets := make([]models.Target, 0)
	existedMapTargets := make(map[string]bool)
	var collector *colly.Collector
	if maxDepth == 0 {
		collector = colly.NewCollector(
			colly.AllowedDomains(domain.Host),
			colly.Async(true),
		)
	} else {
		collector = colly.NewCollector(
			colly.AllowedDomains(domain.Host),
			colly.MaxDepth(maxDepth),
		)
	}
	collector.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 5})
	collector.OnHTML("a[href]", func(element *colly.HTMLElement) {
		absoluteURL := element.Request.AbsoluteURL(element.Attr("href"))
		ok, newTarget, err := c.generateTarget(absoluteURL, domain, existedMapTargets)
		if err != nil {
			log.Printf("Can not generate target from crawled url %s with error %s", absoluteURL, err)
			return
		}
		if !ok {
			return
		}
		targets = append(targets, *newTarget)
		_ = element.Request.Visit(element.Attr("href"))
	})
	collector.Visit(domain.String())
	collector.Wait()
	return targets, nil
}

// generateTarget check if un-match scheme, not include domain host, or duplicate target in existedMapTargets
func (c *CollyCrawler) generateTarget(absoluteURL string, domainHost *url.URL,
	existMapTargets map[string]bool,
) (bool, *models.Target, error) {
	if !strings.Contains(absoluteURL, domainHost.Host) {
		return false, nil, nil
	}
	parsedUrl, err := url.Parse(absoluteURL)
	if err != nil {
		return true, nil, err
	}
	if parsedUrl.Scheme != domainHost.Scheme {
		return false, nil, nil
	}

	hostPathURL := parsedUrl.Scheme + "://" + parsedUrl.Host + parsedUrl.Path
	mapRawQueryURL, _ := url.ParseQuery(parsedUrl.RawQuery)
	querySlice := make([]string, 0)
	for k := range mapRawQueryURL {
		querySlice = append(querySlice, k)
	}
	sort.Strings(querySlice)
	queriesKeyCheckMap := strings.Join(querySlice, ",")
	keyCheckMap := hostPathURL + queriesKeyCheckMap
	if existed := existMapTargets[keyCheckMap]; existed {
		return false, nil, nil
	}
	existMapTargets[keyCheckMap] = true
	return true, &models.Target{
		URL:    parsedUrl,
		Method: http.MethodGet,
		RawURL: parsedUrl.String(),
	}, nil
}
