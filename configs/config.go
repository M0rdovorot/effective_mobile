package configs

import "os"

type DBConfig struct {
	DriverSQL string
	DatabaseDMS string
	DatabaseUser string
	DatabaseUserPassword string
	DatabaseServerIP string
	DatabaseServerPort string
	DatabaseName string
	DatabaseURL string
	MigrationsPath string
	SourceDriver string
}

type Config struct {
	DB DBConfig
	CarInfoAPI string
	BackendServerPort string
}

func New() *Config {
	config := &Config{
		DB: DBConfig{
			DriverSQL: getEnv("DRIVER_SQL", ""),
			DatabaseDMS: getEnv("DATABASE_DMS", ""),
			DatabaseUser: getEnv("DATABASE_USER", ""),
			DatabaseUserPassword: getEnv("DATABASE_USER_PASSWORD", ""),
			DatabaseServerIP: getEnv("DATABASE_SERVER_IP", ""),
			DatabaseServerPort: getEnv("DATABASE_SERVER_PORT", ""),
			DatabaseName: getEnv("DATABASE_NAME", ""),
			MigrationsPath: getEnv("MIGRATIONS_PATH", ""),
			SourceDriver: getEnv("SOURCE_DRIVER", ""),
		},
		BackendServerPort: getEnv("BACKEND_SERVER_PORT", ""),
		CarInfoAPI: getEnv("CAR_INFO_API", ""),
	}
	config.DB.DatabaseURL = config.DB.DatabaseDMS + "://" + config.DB.DatabaseUser +
	":" + config.DB.DatabaseUserPassword + "@" + config.DB.DatabaseServerIP +
	":" + config.DB.DatabaseServerPort + "/" + config.DB.DatabaseName
	return config
}

func getEnv(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultVal
}