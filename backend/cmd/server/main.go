package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// User representa um usuário do sistema
type User struct {
	ID        string
	Email     string
	Password  string
	CreatedAt string
}

// Função para obter variável de ambiente com fallback para valor padrão
func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func getDBConfig() DBConfig {
	return DBConfig{
		Host:     getEnv("DB_HOST", "localhost"),
		Port:     getEnv("DB_PORT", "5432"),
		User:     getEnv("POSTGRES_USER", ""),
		Password: getEnv("POSTGRES_PASSWORD", ""),
		DBName:   getEnv("POSTGRES_DB", ""),
		SSLMode:  getEnv("DB_SSLMODE", "disable"),
	}
}

func connectDB() (*sql.DB, error) {
	if err := godotenv.Load("../.env"); err != nil {
		log.Println("Aviso: Arquivo .env não encontrado, usando variáveis de ambiente do sistema")
	}

	// Obter configurações
	config := getDBConfig()

	// Construir string de conexão
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.User, config.Password, config.DBName, config.SSLMode,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Verificar conexão
	if err = db.Ping(); err != nil {
		db.Close() // Fechar conexão em caso de falha
		return nil, err
	}

	return db, nil
}

// Função para obter todos os usuários
func getAllUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close() // Certifique-se de fechar as linhas após o uso

	var users []User // Cria um slice para armazenar os usuários

	for rows.Next() {
		var user User
		// Lê os dados da linha atual e armazena na variável user
		if err := rows.Scan(&user.ID, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user) // Adiciona o usuário ao slice
	}

	// Verifica se houve erro durante a iteração
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func main() {
	// Conectar ao banco de dados
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Obter todos os usuários
	users, err := getAllUsers(db)
	if err != nil {
		log.Fatalf("Erro ao obter usuários: %v", err)
	}

	// Exibir os usuários
	for _, user := range users {
		fmt.Printf("ID: %s, Email: %s, Created At: %s\n", user.ID, user.Email, user.CreatedAt)
	}
}
