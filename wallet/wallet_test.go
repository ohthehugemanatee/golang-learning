package wallet

import "testing"

// TestWallet tests the wallet struct.
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(wallet, Bitcoin(10), t)
	})
	t.Run("Withdrawal", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Withdraw(Bitcoin(4))
		assertBalance(wallet, Bitcoin(6), t)
	})
	t.Run("Insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(3)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(5))
		assertError(err, errInsufficientFunds, t)
		assertBalance(wallet, startingBalance, t)
	})
}

func assertBalance(wallet Wallet, want Bitcoin, t *testing.T) {
	t.Helper()
	got := wallet.Balance()
	if want != got {
		t.Errorf("Got %s, wanted %s", got, want)
	}
}

func assertError(got error, want error, t *testing.T) {
	t.Helper()
	if got == nil {
		t.Fatal("Expected an error but didn't get one")
	}
	if got != want {
		t.Errorf("Expected error %q but got error %q", want, got)
	}
}
