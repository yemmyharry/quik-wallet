package resource

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	apirouter := router.Group("api/v1")
	apirouter.GET("/wallets/:id/balance", s.GetWalletBalanceByID())
	apirouter.POST("/wallets/:id/credit", s.CreditsWalletByID())
	apirouter.POST("/wallets/:id/debit", s.DebitsWalletByID())
	router.NoRoute(func(c *gin.Context) { c.JSON(404, "no route") })
}
