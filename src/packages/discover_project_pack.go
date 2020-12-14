package packages

type DiscoverProjectPack struct {
	ProjectID  string `json:"project_id"`
	DepthLevel int    `json:"depth_level"` // DepthLevel for limit level crawler crawl url
	IsLoadByJS string `json:"is_load_by_js"`
}
