package resource

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func (s *HTTPHandler) GetWalletBalanceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ID, _ := strconv.Atoi(id)
		//if err != nil {
		//	c.JSON(400, gin.H{
		//		"error": err.Error(),
		//	})
		//	return
		//}

		wallet, err := s.walletService.GetBalance(int64(ID))
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, wallet)
	}
}
