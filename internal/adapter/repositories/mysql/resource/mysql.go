package resource

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

type WalletRepositoryDB struct {
	db *gorm.DB
}

const DB_USERNAME = "root"
const DB_PASSWORD = "12345678"
const DB_NAME = "wallet"
const DB_HOST = "127.0.0.1"
const DB_PORT = "3306"

func NewWalletRepositoryDB() *WalletRepositoryDB {
	dsn := DB_USERNAME + ":" + DB_PASSWORD + "@tcp" + "(" + DB_HOST + ":" + DB_PORT + ")/" + DB_NAME + "?" + "charset=utf8mb4&parseTime=True&loc=Local"
	client, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	//client.AutoMigrate(&domain.Wallet{})
	return &WalletRepositoryDB{client}
}
