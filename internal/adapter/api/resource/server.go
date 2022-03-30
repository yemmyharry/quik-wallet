package resource

import (
	"github.com/gin-gonic/gin"
	ports "quik-wallet/internal/ports/resource"
)

type HTTPHandler struct {
	walletService ports.WalletService
	Route         *gin.Engine
}

func NewHTTPHandler(walletService ports.WalletService, router *gin.Engine) *HTTPHandler {
	handler := &HTTPHandler{
		walletService: walletService,
		Route:         router,
	}
	return handler
}
