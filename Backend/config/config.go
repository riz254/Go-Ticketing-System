package config

import (
	"log"

	"github.com/caarlos0/env"
	"github.com/joho/godotenv"
)

// EnvConfig holds the environment variables required for the application
type EnvConfig struct {
	ServerPort string `env:"SERVER_PORT,required"`
	DBHost     string `env:"DB_HOST,required"`
	DBName     string `env:"DB_NAME,required"` // Use uppercase for the environment variable key
	DBUser     string `env:"DB_USER,required"`
	DBPassword string `env:"DB_PASSWORD,required"`
	DBSSLMode  string `env:"DB_SSLMODE,required"`
}

// NewEnvConfig loads the environment variables and returns the EnvConfig struct
func NewEnvConfig() *EnvConfig {
	// Load .env file into the environment
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Unable to load the .env file: %v", err) // Log but don't crash
	}

	// Parse environment variables into the EnvConfig struct
	config := &EnvConfig{}

	if err := env.Parse(config); err != nil {
		log.Fatalf("Unable to parse environment variables: %v", err)
	}

	return config
}
