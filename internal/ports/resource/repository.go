package ports

import domain "quik-wallet/internal/core/domain/resource"

type WalletRepository interface {
	GetBalance(id int64) (domain.Wallet, error)
	Credit(id int64, amount string) (domain.Wallet, error)
	//Debit(id int64, amount string) (domain.Wallet, error)
}
