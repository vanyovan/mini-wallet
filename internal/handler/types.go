package handler

type InitWalletRequest struct {
	CustomerXid string `json:"customer_xid"`
}

type InitUserResponseData struct {
	Token string `json:"token"`
}
