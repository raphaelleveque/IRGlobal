package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/currency"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

type AppContainer struct {
	userService        domain.UserService
	userHandler        *user.UserHandler
	authService        domain.AuthService
	authHandler        *auth.AuthHandler
	transactionService domain.TransactionService
	transactionHandler *transaction.TransactionHandler
	currencyService domain.CurrencyService
}

func NewAppContainer(db *sql.DB, secretKey []byte) *AppContainer {
	// Repositories
	userRepo := user.NewUserRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db) // Será necessário quando criar

	// Services
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(secretKey, userService)
	currencyService := currency.NewCurrencyService()
	transactionService := transaction.NewTransactionService(transactionRepo, currencyService) // Será necessário quando criar

	// Handlers
	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(userService, authService)
	transactionHandler := transaction.NewTransactionHandler(transactionService) // Será necessário quando criar

	return &AppContainer{
		userService: userService,
		userHandler: userHandler,
		authService: authService,
		authHandler: authHandler,
		transactionService: transactionService, 
		transactionHandler: transactionHandler, 
		currencyService: currencyService,
	}
}

func (c *AppContainer) GetUserHandler() *user.UserHandler {
	return c.userHandler
}

func (c *AppContainer) GetAuthHandler() *auth.AuthHandler {
	return c.authHandler
}

func (c *AppContainer) GetAuthService() domain.AuthService {
	return c.authService
}

func (c *AppContainer) GetTransactionHandler() *transaction.TransactionHandler {
	return c.transactionHandler
}

