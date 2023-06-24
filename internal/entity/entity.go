package entity

import "time"

type UserRequestParam struct {
	CustomerXid string `json:"customer_xid"`
}

type UserResponse struct {
	Token string `json:"token"`
}

type User struct {
	CustomerXid string `json:"customer_xid" db:"user_id"`
	Token       string `json:"token" db:"token"`
}

type Wallet struct {
	WalletId   string     `json:"wallet_id"`
	OwnedBy    string     `json:"owned_by"`
	Status     string     `json:"status"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	Balance    float64    `json:"balance"`
}

type TransactionRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceId string  `json:"reference_id"`
}

type DepositResponse struct {
	Id          string     `json:"id"`
	DepositedBy string     `json:"deposit_by"`
	Status      string     `json:"status"`
	DepositedAt *time.Time `json:"deposited_at"`
	Amount      float64    `json:"amount"`
	ReferenceId string     `json:"reference_id"`
}

type Transaction struct {
	WalletId      string     `json:"wallet_id" db:"wallet_id"`
	TransactionId string     `json:"transaction_id" db:"transaction_id"`
	Status        string     `json:"status" db:"status"`
	TransactedAt  *time.Time `json:"transacted_at" db:"transacted_at"`
	Type          string     `json:"type" db:"type"`
	Amount        float64    `json:"amount" db:"amount"`
	ReferenceId   string     `json:"reference_id" db:"reference_id"`
}
