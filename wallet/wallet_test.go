package wallet

import "testing"

// TestWallet tests the wallet struct.
func TestWallet(t *testing.T) {
	checkBalance := func(wallet Wallet, want Bitcoin, t *testing.T) {
		t.Helper()
		got := wallet.Balance()
		if want != got {
			t.Errorf("Got %s, wanted %s", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		checkBalance(wallet, Bitcoin(10), t)
	})
	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(4))
		checkBalance(wallet, Bitcoin(6), t)
	})
}
