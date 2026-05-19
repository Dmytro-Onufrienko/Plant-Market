package config

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	DBUrl          string
	Port           string
	Env            string
	JWTSecret      string
	JWTExpiry      time.Duration
	AllowedOrigins string
	UploadsDir     string
}

func Load() (*Config, error) {
	// не фатально — змінні можуть бути вже в оточенні (prod/stage)
	_ = godotenv.Load(".env")

	expiryHours, err := strconv.Atoi(getEnv("JWT_EXPIRY_HOURS", "24"))
	if err != nil {
		expiryHours = 24
	}

	cfg := &Config{
		DBUrl:          requireEnv("DATABASE_URL"),
		Port:           getEnv("PORT", "8080"),
		Env:            getEnv("ENV", "local"),
		JWTSecret:      requireEnv("JWT_SECRET"),
		JWTExpiry:      time.Duration(expiryHours) * time.Hour,
		AllowedOrigins: getEnv("ALLOWED_ORIGINS", "http://localhost:5173"),
		UploadsDir:     getEnv("UPLOADS_DIR", "./static/uploads"),
	}

	return cfg, nil
}

func (c *Config) IsProd() bool {
	return c.Env == "prod"
}

func requireEnv(key string) string {
	v := os.Getenv(key)
	if v == "" {
		panic(fmt.Sprintf("required environment variable %q is not set", key))
	}
	return v
}

func getEnv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
