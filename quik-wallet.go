package main

import (
	"github.com/gin-gonic/gin"
	adapter "quik-wallet/internal/adapter/api/resource"
	"quik-wallet/internal/adapter/repositories/mysql/resource"
	"quik-wallet/internal/core/logger"
	services "quik-wallet/internal/core/services/resource"
)

func main() {
	router := gin.Default()
	database := resource.NewWalletRepositoryDB()
	service := services.New(database)
	handler := adapter.NewHTTPHandler(service)
	handler.Routes(router)
	logger.Info("Starting server on port 8080")
	router.Run(":8080")
}
