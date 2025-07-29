package server

import (
	"github.com/leandrodam/transactions/internal/infrastructure"
)

func (s *Server) registerRoutes(handlers infrastructure.Handlers) {
	v1 := s.httpServer.Group("/v1")

	accounts := v1.Group("/accounts")
	accounts.POST("", handlers.Account.Create)
	accounts.GET("/:accountId", handlers.Account.GetByID)

	transactions := v1.Group("/transactions")
	transactions.POST("", handlers.Transaction.Create)
}
