package wallet

import "testing"

// TestWallet tests the wallet struct.
func TestWallet(t *testing.T) {
	wallet := Wallet{}
	wallet.Deposit(10)
	got := wallet.Balance()
	want := 10
	if got != want {
		t.Errorf("Got %d, wanted %d", got, want)
	}
}
