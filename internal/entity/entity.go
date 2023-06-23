package entity

type WalletRequestParam struct {
	CustomerXid string `json:"customer_xid"`
}

type WalletResponse struct {
	Token string `json:"token"`
}

type User struct {
	CustomerXid string `json:"customer_xid" db:"user_id"`
	Token       string `json:"Token" db:"token"`
}
