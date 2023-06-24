package helper

type key string

var userKey = key("user_context")
var TxKey = key("tx_context")

var (
	ConstantEnabled    = "enabled"
	ConstantDisabled   = "disabled"
	ConstantDefaultInt = "0"
	ConstantTimeParsed = "2023-03-06 00:00:00"
	ConstantSuccess    = "success"
	ConstantFailed     = "failed"
	ConstantDeposit    = "deposit"
	ConstantWithdrawal = "withdrawal"
)
