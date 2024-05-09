package configs

const (
	BackendServerPort  = ":8001"

	DriverSQL            = "pgx"
	DatabaseDMS          = "postgres"
	DatabaseUser         = "banners"
	DatabaseUserPassword = "12345"
	// DatabaseServerIP     = "localhost"
	DatabaseServerIP = "effective_mobile-banners-postgres-1"
	DatabaseServerPort   = "5432"
	DatabaseName         = "banners"
	DatabaseURL          = DatabaseDMS + "://" + DatabaseUser +
		":" + DatabaseUserPassword + "@" + DatabaseServerIP +
		":" + DatabaseServerPort + "/" + DatabaseName
	MigrationsPath      = "db/migrations"
	SourceDriver        = "file://"
	BannerTable           = "public.banner"
	TagToBannerTable = "public.banner_to_tag"

	// RedisServerIP   = "127.0.0.1"
	RedisServerPort = "6379"
	RedisServerIP = "effective_mobile-banners-redis-1"

	APIURL = "http://localhost:8001/api/v1/"
)