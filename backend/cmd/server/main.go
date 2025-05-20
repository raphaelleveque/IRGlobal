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
	// Carregar configurações do ambiente
	cfg := config.LoadConfig()

	// Inicializar conexão com o banco de dados
	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer db.Close()

	// Inicializar o container de dependências
	appContainer := container.NewAppContainer(db)

	// Configurar rotas da aplicação
	router := router.SetupRoutes(appContainer.GetUserHandler())

	// Iniciar o servidor HTTP
	serverAddr := ":" + cfg.Port
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
