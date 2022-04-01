package resource

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"net/http"
	"net/http/httptest"
	"quik-wallet/internal/core/domain/resource"
	services "quik-wallet/internal/core/services/mock"
	"strings"
	"testing"
	"time"
)

func TestApplication_GetWalletByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedService := services.NewMockWalletService(ctrl)
	router := gin.Default()

	handler := NewHTTPHandler(mockedService)
	handler.Routes(router)

	t.Run("Should return an error", func(t *testing.T) {
		mockedService.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{}, errors.New("an error occurred"))
		req, err := http.NewRequest(http.MethodGet, "/api/v1/wallets/1/balance", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "an error occurred") {
			t.Fatalf("Expected error message: %s to be in the response body", "an error occurred")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

	t.Run("Get balance by id", func(t *testing.T) {
		mockedService.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{ID: 1, Balance: "100", CreatedAt: time.Now()}, nil)
		req, err := http.NewRequest(http.MethodGet, "/api/v1/wallets/1/balance", nil)
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()

		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "100") {
			t.Fatalf("Expected balance: %s to be in the response body", "100")
		}
		if !strings.Contains(resp.Body.String(), "1") {
			t.Fatalf("Expected balance: %s to be in the response body", "100")
		}
		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusOK, resp.Code)
		}
	})
}

func TestApplication_CreditWalletByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedService := services.NewMockWalletService(ctrl)
	router := gin.Default()

	handler := NewHTTPHandler(mockedService)
	handler.Routes(router)

	t.Run("Should credit wallet", func(t *testing.T) {
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "100",
		}
		mockedService.EXPECT().Credit(int64(1), wallet.Amount).Return(resource.Wallet{ID: 1, Balance: "100", CreatedAt: time.Now()}, nil)
		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/credit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "100") {
			t.Fatalf("Expected balance: %s to be in the response body", "100")
		}
		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("Amount must be greater than zero", func(t *testing.T) {
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "-10",
		}

		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/credit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "Amount must be greater than zero") {
			t.Fatalf("Expected error message: %s to be in the response body", "Amount must be greater than zero")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

	t.Run("Should return an error when it fails to credit wallet", func(t *testing.T) {
		mockedService.EXPECT().Credit(int64(1), "100").Return(resource.Wallet{}, errors.New("an error occurred"))
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "100",
		}
		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/credit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "an error occurred") {
			t.Fatalf("Expected error message: %s to be in the response body", "an error occurred")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})
}

func TestApplication_DebitWalletByID(t *testing.T) {
	ctrl := gomock.NewController(t)

	mockedService := services.NewMockWalletService(ctrl)
	router := gin.Default()

	handler := NewHTTPHandler(mockedService)
	handler.Routes(router)

	t.Run("Amount must be greater than zero", func(t *testing.T) {
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "-10",
		}

		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "Amount must be greater than zero") {
			t.Fatalf("Expected error message: %s to be in the response body", "Amount must be greater than zero")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

	t.Run("Should check if balance is less than amount", func(t *testing.T) {
		mockedService.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{ID: 1, Balance: "100", CreatedAt: time.Now()}, nil)
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "2000",
		}
		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "Your wallet balance has insufficient funds") {
			t.Fatalf("Expected error message: %s to be in the response body", "Your wallet balance has insufficient funds")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

	t.Run("Should debit wallet", func(t *testing.T) {
		mockedService.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{ID: 1, Balance: "100", CreatedAt: time.Now()}, nil)

		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "50",
		}
		mockedService.EXPECT().Debit(int64(1), wallet.Amount).Return(resource.Wallet{ID: 1, Balance: "50", CreatedAt: time.Now()}, nil)
		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		fmt.Println(resp.Body.String(), "bodddy")
		if !strings.Contains(resp.Body.String(), "50") {
			t.Fatalf("Expected balance: %s to be in the response body", "50")
		}
		if resp.Code != http.StatusOK {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusOK, resp.Code)
		}
	})

	t.Run("Should return an error when it fails to debit wallet", func(t *testing.T) {
		mockedService.EXPECT().GetBalance(int64(1)).Return(resource.Wallet{ID: 1, Balance: "100", CreatedAt: time.Now()}, nil)
		mockedService.EXPECT().Debit(int64(1), "100").Return(resource.Wallet{}, errors.New("an error occurred"))
		wallet := struct {
			Amount string `json:"amount"`
		}{
			Amount: "100",
		}
		m, _ := json.Marshal(wallet)
		req, err := http.NewRequest(http.MethodPost, "/api/v1/wallets/1/debit", strings.NewReader(string(m)))
		if err != nil {
			t.Fatal(err)
		}
		resp := httptest.NewRecorder()
		router.ServeHTTP(resp, req)
		if !strings.Contains(resp.Body.String(), "an error occurred") {
			t.Fatalf("Expected error message: %s to be in the response body", "an error occurred")
		}
		if resp.Code != http.StatusBadRequest {
			t.Fatalf("Expected status code: %d but got: %d", http.StatusBadRequest, resp.Code)
		}
	})

}
