package wallet

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC ", b)
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(d Bitcoin) {
	w.balance = w.balance + d
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(b Bitcoin) error {
	if w.balance < b {
		return ErrNoFunds
	}
	w.balance = w.balance - b
	return nil
}

var ErrNoFunds = errors.New("Недостаточно средств, бомж")
