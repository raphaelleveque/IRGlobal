package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config representa a configuração do aplicativo
type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMODE  string
	JWT_Secret string
}

// LoadConfig carrega as configurações a partir das variáveis de ambiente
func LoadConfig() Config {
	if _, err := os.Stat("/.dockerenv"); os.IsNotExist(err) {
		envPath := "../../../.env" 

		if err := godotenv.Load(envPath); err != nil {
			log.Println("No .env file found, using default environment variables")
		}
	}
	return Config{
		Port:       getEnv("BACKEND_PORT", "8080"),
		DBHost:     getDBHost(),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("POSTGRES_USER", ""),
		DBPassword: getEnv("POSTGRES_PASSWORD", ""),
		DBName:     getEnv("POSTGRES_DB", ""),
		DBSSLMODE:  getEnv("DB_SSLMODE", "disable"),
		JWT_Secret: getEnv("JWT_SECRET", ""),
	}
}

// getDBHost detecta se está rodando dentro de um container Docker
func getDBHost() string {
	if _, err := os.Stat("/.dockerenv"); err == nil {
		// Estamos dentro de um container Docker
		return getEnv("DB_HOST", "")
	}
	// Estamos fora do container
	return "localhost"
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
