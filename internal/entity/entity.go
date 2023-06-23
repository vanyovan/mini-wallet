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
