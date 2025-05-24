package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/currency"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/infrastructure"
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
}

func NewAppContainer(db *sql.DB, secretKey []byte) *AppContainer {
	// Repositories
	userRepo := user.NewUserRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)
	positionRepo := position.NewPositionRepository(db)

	// Services
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(secretKey, userService)
	currencyService := currency.NewCurrencyService()

	// Criar positionService primeiro sem transactionService
	positionService := position.NewPositionService(positionRepo, nil)
	// Depois criar transactionService com positionService
	transactionService := transaction.NewTransactionService(transactionRepo, currencyService, positionService)
	// Atualizar positionService com transactionService
	positionService = position.NewPositionService(positionRepo, transactionService)

	// UnitOfWork factory
	uowFactory := func() domain.UnitOfWork {
		return infrastructure.NewUnitOfWork(db)
	}

	// Handlers
	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(userService, authService)
	transactionHandler := transaction.NewTransactionHandler(transactionService, uowFactory)
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
