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
	log.Printf("Starting IRGlobal backend server...")

	cfg := config.LoadConfig()
	log.Printf("Configuration loaded successfully")

	db, err := database.Connect(cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMODE)
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		return
	}
	defer db.Close()

	appContainer := container.NewAppContainer(db)
	log.Printf("Application container initialized")

	router := router.SetupRoutes(appContainer.GetUserHandler())
	log.Printf("Routes configured successfully")

	serverAddr := ":" + cfg.Port
	log.Printf("Server starting on port %s", cfg.Port)
	if err := http.ListenAndServe(serverAddr, router); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
