package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	Server struct {
		Port string
		Host string
	}

	Database struct {
		Host     string
		Port     string
		User     string
		Password string
		DbName   string
		SslMode  string
	}

	JWT struct {
		SecretKey string
		TokenExp  time.Duration
	}

	Env         string
	Date_layout string
}

func Load() (*Config, error) {
	godotenv.Load()

	cfg := &Config{}

	cfg.Server.Host = getEnv("SERVER_HOST", "localhost")
	cfg.Server.Port = getEnv("SERVER_PORT", "8080")

	cfg.Database.Host = getEnv("DB_HOST", "localhost")
	cfg.Database.Port = getEnv("DB_PORT", "5432")
	cfg.Database.User = getEnv("DB_USER", "postgres")
	cfg.Database.Password = getEnv("DB_PASSWORD", "")
	cfg.Database.DbName = getEnv("DB_NAME", "smart_expense")
	cfg.Database.SslMode = getEnv("DB_SSLMODE", "disable")

	cfg.JWT.SecretKey = getEnv("JWT_SECRET", "your-secure-secret-key")
	cfg.JWT.TokenExp = time.Hour * 1

	cfg.Env = getEnv("ENV", "development")
	cfg.Date_layout = getEnv("DATE_LAYOUT", "2006-01-02 15:04:05")

	return cfg, nil
}

func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}
