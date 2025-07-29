package server

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/leandrodam/transactions/internal/infrastructure"
	"github.com/leandrodam/transactions/internal/infrastructure/config"
)

func NewResources(config config.Database) infrastructure.Resources {
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.User, config.Password, config.Host, config.Port, config.Name))
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	return infrastructure.Resources{
		DB: db,
	}
}
