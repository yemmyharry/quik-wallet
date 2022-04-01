package services

import (
	"errors"
	"github.com/golang/mock/gomock"
	"quik-wallet/internal/core/domain/resource"
	services "quik-wallet/internal/core/services/mock"
	"testing"
)

func TestApplication_GetBalance(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepository := services.NewMockWalletRepository(ctrl)
	service := New(mockedRepository)
	t.Run("Test for an error", func(t *testing.T) {
		mockedRepository.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{}, errors.New("an error occurred"))

		_, err := service.GetBalance(int64(1))
		if err.Error() != "an error occurred" {
			t.Errorf("Expected error: %s, got: %s", "an error occurred", err.Error())
		}
	})
	t.Run("Credit wallet service", func(t *testing.T) {
		mockedRepository.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{ID: 1, Balance: "100"}, nil)

		resp, err := service.GetBalance(int64(1))
		if err != nil {
			t.Errorf("Expected error: %v, got: %s", nil, err.Error())
		}
		if resp.Balance != "100" {
			t.Errorf("Expected balance: %s, got: %s", "100", resp.Balance)
		}
		if resp.ID != 1 {
			t.Errorf("Expected id: %d, got: %d", 1, resp.ID)
		}
	})
}

func TestApplication_Credit(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepository := services.NewMockWalletRepository(ctrl)
	service := New(mockedRepository)
	t.Run("Test for an error", func(t *testing.T) {
		mockedRepository.EXPECT().Credit(int64(1), "100").Return(resource.Wallet{}, errors.New("an error occurred"))

		_, err := service.Credit(int64(1), "100")
		if err.Error() != "an error occurred" {
			t.Errorf("Expected error: %s, got: %s", "an error occurred", err.Error())
		}
	})
	t.Run("Credit wallet service", func(t *testing.T) {
		mockedRepository.EXPECT().Credit(int64(1), "100").Return(resource.Wallet{ID: 1, Balance: "100"}, nil)

		resp, err := service.Credit(int64(1), "100")
		if err != nil {
			t.Errorf("Expected error: %v, got: %s", nil, err.Error())
		}
		if resp.Balance != "100" {
			t.Errorf("Expected balance: %s, got: %s", "100", resp.Balance)
		}
		if resp.ID != 1 {
			t.Errorf("Expected id: %d, got: %d", 1, resp.ID)
		}
	})
}

func TestApplication_Debit(t *testing.T) {
	ctrl := gomock.NewController(t)
	mockedRepository := services.NewMockWalletRepository(ctrl)
	service := New(mockedRepository)
	t.Run("Test for an error", func(t *testing.T) {
		mockedRepository.EXPECT().Debit(int64(1), "100").Return(resource.Wallet{}, errors.New("an error occurred"))

		_, err := service.Debit(int64(1), "100")
		if err.Error() != "an error occurred" {
			t.Errorf("Expected error: %s, got: %s", "an error occurred", err.Error())
		}
	})
	t.Run("Debit wallet service", func(t *testing.T) {
		mockedRepository.EXPECT().Debit(int64(1), "100").Return(resource.Wallet{ID: 1, Balance: "100"}, nil)

		resp, err := service.Debit(int64(1), "100")
		if err != nil {
			t.Errorf("Expected error: %v, got: %s", nil, err.Error())
		}
		if resp.Balance != "100" {
			t.Errorf("Expected balance: %s, got: %s", "100", resp.Balance)
		}
		if resp.ID != 1 {
			t.Errorf("Expected id: %d, got: %d", 1, resp.ID)
		}
	})
}
