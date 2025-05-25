package container

import (
	"database/sql"

	"github.com/raphaelleveque/IRGlobal/backend/internal/auth"
	"github.com/raphaelleveque/IRGlobal/backend/internal/currency"
	"github.com/raphaelleveque/IRGlobal/backend/internal/domain"
	"github.com/raphaelleveque/IRGlobal/backend/internal/infrastructure"
	"github.com/raphaelleveque/IRGlobal/backend/internal/position"
	realizedpnl "github.com/raphaelleveque/IRGlobal/backend/internal/realized_pnl"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction"
	"github.com/raphaelleveque/IRGlobal/backend/internal/transaction/orchestrator"
	"github.com/raphaelleveque/IRGlobal/backend/internal/user"
)

type AppContainer struct {
	userService        domain.UserService
	userHandler        *user.UserHandler
	authService        domain.AuthService
	authHandler        *auth.AuthHandler
	transactionService domain.TransactionService
	transactionHandler *transaction.TransactionHandler
	currencyService    domain.CurrencyService
	positionService    domain.PositionService
	positionHandler    *position.PositionHandler
	realizedPNLService domain.RealizedPNLService
	realizedPNLHandler *realizedpnl.RealizedPNLHandler
	unitOfWork         domain.UnitOfWork
	orchestrator       *orchestrator.TransactionOrchestrator
}

func NewAppContainer(db *sql.DB, secretKey []byte) *AppContainer {
	// Infrastructure
	unitOfWork := infrastructure.NewSqlUnitOfWork(db)

	// Repositories
	userRepo := user.NewUserRepository(db)
	transactionRepo := transaction.NewTransactionRepository(db)
	positionRepo := position.NewPositionRepository(db)
	realizedpnlRepo := realizedpnl.NewRealizedPNLRepository(db)

	// Services
	userService := user.NewUserService(userRepo)
	authService := auth.NewAuthService(secretKey, userService)
	currencyService := currency.NewCurrencyService()
	transactionService := transaction.NewTransactionService(transactionRepo, currencyService)
	positionService := position.NewPositionService(positionRepo, transactionService)
	realizedpnlService := realizedpnl.NewRealizedPNLService(realizedpnlRepo)

	// Orchestrator
	transactionOrchestrator := orchestrator.NewTransactionOrchestrator(transactionService, positionService, realizedpnlService, unitOfWork)

	// Handlers
	userHandler := user.NewUserHandler(userService)
	authHandler := auth.NewAuthHandler(userService, authService)
	transactionHandler := transaction.NewTransactionHandler(transactionService, transactionOrchestrator)
	positionHandler := position.NewPositionHandler(positionService)
	realizedPNLHandler := realizedpnl.NewRealizedPNLHandler(realizedpnlService)

	return &AppContainer{
		userService:        userService,
		userHandler:        userHandler,
		authService:        authService,
		authHandler:        authHandler,
		transactionService: transactionService,
		transactionHandler: transactionHandler,
		currencyService:    currencyService,
		positionService:    positionService,
		positionHandler:    positionHandler,
		realizedPNLService: realizedpnlService,
		realizedPNLHandler: realizedPNLHandler,
		unitOfWork:         unitOfWork,
		orchestrator:       transactionOrchestrator,
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
