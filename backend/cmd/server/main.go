package main

import (
	"log"
	"net/http"

	"github.com/raphaelleveque/IRGlobal/backend/internal/container"
	"github.com/raphaelleveque/IRGlobal/backend/internal/database"
	"github.com/raphaelleveque/IRGlobal/backend/internal/router"
	"github.com/raphaelleveque/IRGlobal/backend/pkg/config"

	_ "github.com/lib/pq"
)

func main() {
	// Carregar configurações
	cfg := config.LoadConfig()

	// Conectar ao banco de dados
	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicializar o container de dependências
	container := container.NewContainer(db)

	// Configurar rotas
	router := router.SetupRoutes(container.UserHandler)

	// Iniciar o servidor
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := http.ListenAndServe(":"+cfg.Port, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
