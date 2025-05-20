package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq"
	"github.com/raphaelleveque/IRGlobal/backend/internal/container"
	"github.com/raphaelleveque/IRGlobal/backend/internal/database"
	"github.com/raphaelleveque/IRGlobal/backend/internal/router"
	"github.com/raphaelleveque/IRGlobal/backend/pkg/config"
)

func main() {
	cfg := config.LoadConfig()
	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)
	if err != nil {
		log.Fatal("Conexãon com o banco de dados não concluída")
		return
	}
	defer db.Close()

	appContainer := container.NewAppContainer(db)
	router := router.SetupRoutes(appContainer.GetUserHandler())

	serverAddr := ":" + cfg.Port
	log.Printf("Servidor iniciado na porta %s", cfg.Port)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
