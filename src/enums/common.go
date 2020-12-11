package enums

import "regexp"

var ResultExistVul = 2
var ResultNotExistVul = 1

var TypeStaticSite = "static"
var TypeDynamicSite = "dynamic"

var DefaultMaxDepth = 2

var EmailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
