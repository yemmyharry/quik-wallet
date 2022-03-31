package resource

import (
	"quik-wallet/internal/core/domain/resource"
	"quik-wallet/internal/core/logger"
)

func (r *WalletRepositoryDB) GetBalance(id int64) (resource.Wallet, error) {
	var balance resource.Wallet
	err := r.db.First(&balance, id)
	if err != nil {
		logger.Error(err.Error)
		return balance, err.Error
	}
	return balance, err.Error
}
