package packages

type DiscoverProjectPack struct {
	ProjectID  string `json:"project_id"`
	DeepLevel  int    `json:"deep_level"` // DeepLevel for limit level crawler crawl url
	IsLoadByJS string `json:"is_load_by_js"`
}
