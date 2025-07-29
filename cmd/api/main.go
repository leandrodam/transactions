package main

import (
	"github.com/labstack/echo/v4"
	"github.com/leandrodam/transactions/internal/infrastructure/server"
)

func main() {
	server.NewServer(echo.New()).Start()
}
