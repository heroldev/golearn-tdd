package pointers

import (
	"errors"
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

var ErrInsufficientFunds = errors.New("cannot withdraw insufficient funds")

func (w *Wallet) Withdraw(amount Bitcoin) error {

	if w.Balance() < amount {
		return ErrInsufficientFunds
	}

	w.balance -= amount

	return nil
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func main() {

}
