package crawler

type baseCrawler struct {
	maximumTarget int
}

func NewBaseCrawlerDriver(maximumTarget int) *baseCrawler {
	return &baseCrawler{
		maximumTarget: maximumTarget,
	}
}
