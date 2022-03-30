package resource

import "quik-wallet/internal/core/domain/resource"

func (r *WalletRepositoryDB) GetBalance(userId int64) (resource.Wallet, error) {
	var wallet resource.Wallet
	err := r.db.QueryRow("SELECT * FROM wallets WHERE user_id = ?", userId)
	if err != nil {
		return wallet, err.Err()
	}
	return wallet, err.Err()
}
