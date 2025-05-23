package main

import (
	"log"

	_ "github.com/lib/pq"
	_ "github.com/raphaelleveque/IRGlobal/backend/docs" // Importação para o Swagger
	"github.com/raphaelleveque/IRGlobal/backend/internal/container"
	"github.com/raphaelleveque/IRGlobal/backend/internal/database"
	"github.com/raphaelleveque/IRGlobal/backend/internal/router"
	"github.com/raphaelleveque/IRGlobal/backend/pkg/config"
)

// @title           IRGlobal API
// @version         1.0
// @description     API para o sistema IRGlobal
// @host      localhost:8080
// @BasePath  /

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

	secretKey := []byte(cfg.JWT_Secret)

	appContainer := container.NewAppContainer(db, secretKey)
	log.Printf("Application container initialized")

	router := router.SetupRoutes(appContainer)
	log.Printf("Routes configured successfully")

	serverAddr := ":" + cfg.Port
	log.Printf("Server starting on port %s", cfg.Port)
	if err := router.Run(serverAddr); err != nil {
		log.Fatalf("Server failed to start: %v", err)
	}
}
