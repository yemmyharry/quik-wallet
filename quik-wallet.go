package main

import (
	"github.com/gin-gonic/gin"
	adapter "quik-wallet/internal/adapter/api/resource"
	"quik-wallet/internal/adapter/repositories/mysql/resource"
	services "quik-wallet/internal/core/services/resource"
)

func main() {
	//start server using go gin framework
	router := gin.Default()
	database := resource.NewWalletRepositoryDB()
	service := services.New(database)
	handler := adapter.NewHTTPHandler(service, router)
	handler.Routes(router)
	handler.Route.Run(":8080")
	//log.Fatal(router.Run(":8080"))
}
