package configs

type Config struct {
	DBHost           string
	DBPort           string
	DBUser           string
	DBPassword       string
	DBName           string
	SQLMapServerHost string
	SQLMapServerPort string
	AuthSecretKet    string
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfig() {
	Common = &Config{
		DBHost:           GetStringWithDefault("DB_HOST", "0.0.0.0"),
		DBPort:           GetStringWithDefault("DB_PORT", "3306"),
		DBUser:           GetString("DB_USER"),
		DBPassword:       GetString("DB_PASSWORD"),
		DBName:           GetString("DB_NAME"),
		SQLMapServerHost: GetStringWithDefault("SQLMAP_SERVER_HOST", "0.0.0.0"),
		SQLMapServerPort: GetStringWithDefault("SQLMAP_SERVER_PORT", "8775"),
		AuthSecretKet:    GetString("AUTH_SECRET_KEY"),
	}
}
