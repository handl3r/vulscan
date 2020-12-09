package configs

import (
	"os"
	"strconv"
	"time"
)

func GetString(key string) string {
	return os.Getenv(key)
}

func GetStringWithDefault(key string, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
func GetInt(key string, defaultValue int) int {
	strValue := os.Getenv(key)
	if strValue == "" {
		return defaultValue
	}
	value, err := strconv.Atoi(strValue)
	if err != nil {
		return defaultValue
	}
	return value
}

func GetDurationWithDefault(key string, defaultTime time.Duration) time.Duration {
	strValue := os.Getenv(key)
	if len(strValue) == 0 {
		return defaultTime
	}
	unit := strValue[len(strValue)-1]
	numericValue, _ := strconv.ParseInt(strValue[0:len(strValue)-1], 10, 64)
	var duration time.Duration
	switch unit {
	case 'S':
		duration = time.Duration(numericValue) * time.Second
	case 'M':
		duration = time.Duration(numericValue) * time.Minute
	case 'H':
		duration = time.Duration(numericValue) * time.Hour
	case 'D':
		duration = time.Duration(numericValue) * 24 * time.Hour
	}
	return duration
}
