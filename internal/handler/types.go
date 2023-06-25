package handler

type InitWalletRequest struct {
	CustomerXid string `json:"customer_xid"`
}

type TransactionRequest struct {
	Amount      float64 `json:"amount"`
	ReferenceId string  `json:"reference_id"`
}
