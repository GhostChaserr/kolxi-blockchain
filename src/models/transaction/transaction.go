package transaction

import (
	"ghostchaser/kolxi-blockchain/src/models/account"
)

type Transaction struct {
	Id     string `json:"id"`
	From   account.Account
	To     account.Account
	Amount int `json:"amount"`
}
