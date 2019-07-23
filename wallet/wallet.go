package wallet

// Wallet is a virtual wallet.
type Wallet struct {
	balance int
}

// Deposit adds an amount to the balance.
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() int {
	return w.balance
}
