package resource

import "quik-wallet/internal/core/domain/resource"

func (r *WalletRepositoryDB) GetBalance(id int64) (resource.Wallet, error) {
	var balance resource.Wallet
	err := r.db.First(&balance, id)
	if err != nil {
		return balance, err.Error
	}
	return balance, err.Error
}
