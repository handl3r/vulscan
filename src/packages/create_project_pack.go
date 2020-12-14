package packages

import "vulscan/src/helpers"

type CreateProjectPack struct {
	Name   string `json:"name"`
	Domain string `json:"domain"`
}

func ValidateCreateProjectPack(pack *CreateProjectPack) bool {
	if len(pack.Name) == 0 {
		return false
	}
	return helpers.IsValidURL(pack.Domain)
}
