package resource

import (
	"github.com/gin-gonic/gin"
)

func (s *HTTPHandler) Routes(router *gin.Engine) {
	apirouter := router.Group("api")
	apirouter.GET("/wallet/balance", s.GetWalletBalanceByID())
	router.NoRoute(func(c *gin.Context) { c.JSON(404, "no route") })
}
