package crawler_driver

type baseCrawlerDriver struct {
	maximumTarget int
}

func NewBaseCrawlerDriver(maximumTarget int) *baseCrawlerDriver {
	return &baseCrawlerDriver{
		maximumTarget: maximumTarget,
	}
}
