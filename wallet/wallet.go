package wallet

import "fmt"

// Bitcoin is bitcoin.
type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

// Wallet is a virtual wallet.
type Wallet struct {
	balance Bitcoin
}

// Deposit adds an amount to the balance.
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

// Withdraw an amount from the balance.
func (w *Wallet) Withdraw(amount Bitcoin) {
	w.balance -= amount
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
