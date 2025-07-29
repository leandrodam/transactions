package infrastructure

import (
	"database/sql"

	accounthandler "github.com/leandrodam/transactions/internal/adapters/http/handlers/account"
	transactionhandler "github.com/leandrodam/transactions/internal/adapters/http/handlers/transaction"
	accountdomain "github.com/leandrodam/transactions/internal/domain/account"
	transactiondomain "github.com/leandrodam/transactions/internal/domain/transaction"
	accountusecase "github.com/leandrodam/transactions/internal/usecases/account"
	transactionusecase "github.com/leandrodam/transactions/internal/usecases/transaction"
)

type Application struct {
	Resources Resources
	Services  Services
	UseCases  UseCases
	Handlers  Handlers
}

type Resources struct {
	DB *sql.DB
}

type Services struct {
	Account     accountdomain.Repository
	Transaction transactiondomain.Repository
}

type UseCases struct {
	Account     accountusecase.UseCase
	Transaction transactionusecase.UseCase
}

type Handlers struct {
	Account     accounthandler.Handler
	Transaction transactionhandler.Handler
}

// TODO: Remover?
func NewApplication(handlers Handlers) Application {
	return Application{
		Resources: Resources{},
		Services:  Services{},
		UseCases:  UseCases{},
		Handlers:  handlers,
	}
}
