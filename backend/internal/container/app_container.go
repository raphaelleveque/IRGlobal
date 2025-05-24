package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/currency"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/position"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

type AppContainer struct {
	db                 *sql.DB
	userService        domain.UserService
	userHandler        *user.UserHandler
	authService        domain.AuthService
	authHandler        *auth.AuthHandler
	transactionService domain.TransactionService
	transactionHandler *transaction.TransactionHandler
	currencyService    domain.CurrencyService
	positionService    domain.PositionService
	positionHandler    *position.PositionHandler
	coordinator        domain.TransactionCoordinatorService
}

func NewAppContainer(db *sql.DB, secretKey []byte) *AppContainer {
	// Repositories
	userRepo := user.NewUserRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)
	positionRepo := position.NewPositionRepository(db)

	// Services - sem dependência de banco de dados
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(secretKey, userService)
	currencyService := currency.NewCurrencyService()

	// Services de domínio - apenas dependências de outros services e repositories
	transactionService := transaction.NewTransactionService(transactionRepo, currencyService)
	positionService := position.NewPositionService(positionRepo, transactionService)

	// Coordinator que gerencia transações e posições - usa os services para lógica de negócio
	coordinator := transaction.NewTransactionCoordinator(db, transactionService, positionService, currencyService)

	// Handlers
	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(userService, authService)
	transactionHandler := transaction.NewTransactionHandler(coordinator)
	positionHandler := position.NewPositionHandler(positionService)

	return &AppContainer{
		db:                 db,
		userService:        userService,
		userHandler:        userHandler,
		authService:        authService,
		authHandler:        authHandler,
		transactionService: transactionService,
		transactionHandler: transactionHandler,
		currencyService:    currencyService,
		positionService:    positionService,
		positionHandler:    positionHandler,
		coordinator:        coordinator,
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
