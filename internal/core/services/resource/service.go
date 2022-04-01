package services

import (
	"quik-wallet/internal/core/domain/resource"
	ports "quik-wallet/internal/ports/resource"
)

type service struct {
	walletRepository ports.WalletRepository
}

func New(walletRepository ports.WalletRepository) *service {
	return &service{
		walletRepository: walletRepository,
	}
}

func (s *service) GetBalance(walletId int64) (resource.Wallet, error) {
	wallet, err := s.walletRepository.GetBalance(walletId)
	if err != nil {
		return resource.Wallet{}, err
	}
	return wallet, nil
}

func (s *service) Credit(walletId int64, amount string) (resource.Wallet, error) {
	wallet, err := s.walletRepository.Credit(walletId, amount)
	if err != nil {
		return resource.Wallet{}, err
	}
	return wallet, nil
}
