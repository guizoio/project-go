package infrastructure

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Environment string
	App         string
	DbHost      string
	DbPort      string
	DbUser      string
	DbPassword  string
	DbDatabase  string
}

func NewConfig() Config {
	if os.Getenv("ENVIRONMENT") == "" {
		if err := godotenv.Load(".env"); err != nil {
			log.Fatalln("Error loading env file")
		}
	}

	return Config{
		Environment: os.Getenv("ENVIRONMENT"),
		App:         os.Getenv("APP"),
		DbHost:      os.Getenv("DB_HOST_LOCAL"),
		DbPort:      os.Getenv("DB_PORT"),
		DbUser:      os.Getenv("DB_USERNAME"),
		DbPassword:  os.Getenv("DB_PASSWORD"),
		DbDatabase:  os.Getenv("DB_DATABASE"),
	}
}

func (c Config) IsLocal() bool {
	return c.Environment == "local"
}

func (c Config) IsProduction() bool {
	return c.Environment == "production"
}

func (c Config) IsDevelopment() bool {
	return c.Environment == "development"
}
