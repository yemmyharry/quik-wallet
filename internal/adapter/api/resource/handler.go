package resource

import (
	"github.com/gin-gonic/gin"
	"quik-wallet/internal/core/logger"
	"strconv"
)

func (s *HTTPHandler) GetWalletBalanceByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ID, _ := strconv.Atoi(id)

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

func (s *HTTPHandler) CreditsWalletByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ID, _ := strconv.Atoi(id)
		w2 := struct {
			Amount string `json:"amount"`
		}{}
		err := c.ShouldBindJSON(&w2)
		if err != nil {
			logger.Error(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		wallet, err := s.walletService.Credit(int64(ID), w2.Amount)
		if err != nil {
			logger.Error(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, wallet)
	}
}

func (s *HTTPHandler) DebitsWalletByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		id := c.Param("id")
		ID, _ := strconv.Atoi(id)
		w2 := struct {
			Amount string `json:"amount"`
		}{}
		err := c.ShouldBindJSON(&w2)
		if err != nil {
			logger.Error(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		wallet, err := s.walletService.Debit(int64(ID), w2.Amount)
		if err != nil {
			logger.Error(err)
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}

		c.JSON(200, wallet)
	}
}
