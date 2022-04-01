package resource

import (
	ports "quik-wallet/internal/ports/resource"
)

type HTTPHandler struct {
	walletService ports.WalletService
}

func NewHTTPHandler(walletService ports.WalletService) *HTTPHandler {
	handler := &HTTPHandler{
		walletService: walletService,
	}
	return handler
}
