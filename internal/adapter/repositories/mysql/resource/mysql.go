package resource

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

type WalletRepositoryDB struct {
	db *sql.DB
}

func NewWalletRepositoryDB() *WalletRepositoryDB {
	client, err := sql.Open("mysql", "root:12345678@tcp(localhost:3306)/wallet")
	if err != nil {
		log.Fatal(err)
	}
	return &WalletRepositoryDB{client}
}
