package repo

import (
	"database/sql"
)

type TransactionRepo interface {
}

func NewTransactionRepo(db *sql.DB) WalletRepo {
	return &Repo{
		db: db,
	}
}
