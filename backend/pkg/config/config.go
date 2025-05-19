package config

import (
	"os"
)

// Config representa a configuração do aplicativo
type Config struct {
	Port       string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMODE string
}

// LoadConfig carrega as configurações a partir das variáveis de ambiente
func LoadConfig() Config {
	return Config{
		Port:       getEnv("BACKEND_PORT", "8080"),
		DBHost:     getEnv("DB_HOST", "localhost"),
		DBPort:     getEnv("DB_PORT", "5432"),
		DBUser:     getEnv("POSTGRES_USER", ""),
		DBPassword: getEnv("POSTGRES_PASSWORD", ""),
		DBName:     getEnv("POSTGRES_DB", ""),
		DBSSLMODE:  getEnv("DB_SSLMODE", "disable"),
	}
}

// getEnv obtém o valor de uma variável de ambiente ou retorna um valor padrão
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
