package helpers

import "net/url"

func IsValidURL(str string) bool {
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return false
	}
	u, err := url.Parse(str)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}
