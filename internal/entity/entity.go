package entity

import "time"

// DTO from handler layer to usecase layer

// DTO from user API
type UserRequestParam struct {
	CustomerXid string `json:"customer_xid"`
}

// init wallet response
type UserResponse struct {
	Token string `json:"token"`
}

// store user object
type User struct {
	CustomerXid string `json:"customer_xid" db:"user_id"`
	Token       string `json:"token" db:"token"`
}

// store wallet object
type Wallet struct {
	WalletId   string     `json:"wallet_id"`
	OwnedBy    string     `json:"owned_by"`
	Status     string     `json:"status"`
	EnabledAt  *time.Time `json:"enabled_at,omitempty"`
	DisabledAt *time.Time `json:"disabled_at,omitempty"`
	Balance    float64    `json:"balance"`
}

// request for deposit or withdraw
type TransactionRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceId string  `json:"reference_id"`
}

// deposit API response
type DepositResponse struct {
	Id          string     `json:"id"`
	DepositedBy string     `json:"deposit_by"`
	Status      string     `json:"status"`
	DepositedAt *time.Time `json:"deposited_at"`
	Amount      float64    `json:"amount"`
	ReferenceId string     `json:"reference_id"`
}

// withdrawn API response
type WithdrawalResponse struct {
	Id          string     `json:"id"`
	WithdrawBy  string     `json:"withdrawn_by"`
	Status      string     `json:"status"`
	WithdrawAt  *time.Time `json:"withdrawn_at"`
	Amount      float64    `json:"amount"`
	ReferenceId string     `json:"reference_id"`
}

// view transaction API response
type Transaction struct {
	WalletId      string     `json:"wallet_id" db:"wallet_id"`
	TransactionId string     `json:"transaction_id" db:"transaction_id"`
	Status        string     `json:"status" db:"status"`
	TransactedAt  *time.Time `json:"transacted_at" db:"transacted_at"`
	Type          string     `json:"type" db:"type"`
	Amount        float64    `json:"amount" db:"amount"`
	ReferenceId   string     `json:"reference_id" db:"reference_id"`
}
