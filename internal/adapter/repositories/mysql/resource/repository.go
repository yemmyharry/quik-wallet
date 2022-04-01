package resource

import (
	"github.com/shopspring/decimal"
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

func (r *WalletRepositoryDB) Credit(id int64, amount string) (resource.Wallet, error) {
	var wallet resource.Wallet
	_ = r.db.Where("id = ?", id).First(&wallet)
	aDec, _ := decimal.NewFromString(amount)
	bDec, _ := decimal.NewFromString(wallet.Balance)
	wallet.Balance = aDec.Add(bDec).String()
	err := r.db.Save(&wallet)
	if err != nil {
		logger.Error(err.Error)
		return wallet, err.Error
	}
	return wallet, nil
}
