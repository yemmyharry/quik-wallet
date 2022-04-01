package main

import (
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
	adapter "quik-wallet/internal/adapter/api/resource"
	"quik-wallet/internal/adapter/repositories/mysql/resource"
	"quik-wallet/internal/core/logger"
	services "quik-wallet/internal/core/services/resource"
)

func main() {
	err := godotenv.Load("quik-wallet.env")
	if err != nil {
		logger.Error("Error loading .env file")
	}
	router := gin.Default()
	database := resource.NewWalletRepositoryDB()
	service := services.New(database)
	handler := adapter.NewHTTPHandler(service)
	handler.Routes(router)
	logger.Info("Starting server on port 8080")
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router.Run(":" + port)
}
