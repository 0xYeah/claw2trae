package config

import "os"

const (
	ProjectName     = "claw2trae"
	ProjectVersion  = "v0.0.2"
	ProjectBundleID = "com.0xYeah.claw2trae"
)

type Config struct {
	Port   string
	DBPath string
}

func Load() *Config {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbPath := os.Getenv("DB_PATH")
	if dbPath == "" {
		dbPath = "./wms.db"
	}

	return &Config{
		Port:   port,
		DBPath: dbPath,
	}
}
