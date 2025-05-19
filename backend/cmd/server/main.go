package main

import (
	"log"
	"net/http"
	"github.com/raphaelleveque/IRGlobal/internal/database"
	"github.com/raphaelleveque/IRGlobal/pkg/config"

	_ "github.com/lib/pq"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Iniciar o servidor
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, nil); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
