package wallet

import "testing"

// TestWallet tests the wallet struct.
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)
		if got != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	})
	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(4))
		got := wallet.Balance()
		want := Bitcoin(6)
		if got != want {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	})
}
