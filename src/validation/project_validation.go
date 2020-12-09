package validation

import (
	"vulscan/src/helpers"
	"vulscan/src/packages"
)

func ValidateCreateProjectPack(pack *packages.CreateProjectPack) bool {
	if len(pack.Name) == 0 {
		return false
	}
	return helpers.IsValidURL(pack.Domain)
}
