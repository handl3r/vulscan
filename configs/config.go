package configs

type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var Common *Config

func Get() *Config {
	return Common
}

func LoadConfig() {
	Common = &Config{
		DBHost:     GetStringWithDefault("DB_HOST", "0.0.0.0"),
		DBPort:     GetStringWithDefault("DB_PORT", "3306"),
		DBUser:     GetString("DB_USER"),
		DBPassword: GetString("DB_PASSWORD"),
		DBName:     GetString("DB_NAME"),
	}
}
