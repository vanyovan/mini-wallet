package handler

type InitWalletRequest struct {
	CustomerXid string `json:"customer_xid"`
}

type InitWalletResponseData struct {
	Token string `json:"token"`
}
