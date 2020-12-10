package packages

type DiscoverProjectPack struct {
	ProjectID  string `json:"project_id"`
	DeepLevel  int    `json:"deep_level"` // DeepLevel for limit level crawler crawl url
	IsLoadByJS int    `json:"is_load_by_js"`
}
