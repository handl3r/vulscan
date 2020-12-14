package configs

import "time"

type Config struct {
	ServerAddress        string
	DBHost               string
	DBPort               string
	DBUser               string
	DBPassword           string
	DBName               string
	SQLMapServerHost     string
	SQLMapServerPort     string
	AuthSecretKet        string
	AccessTokenExp       time.Duration
	MaximumTargetCrawler int
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfig() {
	loadEnvironments()
	Common = &Config{
		ServerAddress:        GetString("SERVER_ADDRESS"),
		DBHost:               GetStringWithDefault("DB_HOST", "0.0.0.0"),
		DBPort:               GetStringWithDefault("DB_PORT", "3306"),
		DBUser:               GetString("DB_USER"),
		DBPassword:           GetString("DB_PASSWORD"),
		DBName:               GetString("DB_NAME"),
		SQLMapServerHost:     GetStringWithDefault("SQLMAP_SERVER_HOST", "http://0.0.0.0"),
		SQLMapServerPort:     GetStringWithDefault("SQLMAP_SERVER_PORT", "8775"),
		AuthSecretKet:        GetString("AUTH_SECRET_KEY"),
		AccessTokenExp:       GetDurationWithDefault("ACCESS_EXP_TIME", 60*24*time.Hour),
		MaximumTargetCrawler: GetInt("MAXIMUM_TARGET_CRAWLER", 10),
	}
}
