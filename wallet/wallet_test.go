package wallet

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		want := Bitcoin(10)

		wallet.Deposit(10)

		got := wallet.Balance()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		want := Bitcoin(10)

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		if want != got {
			t.Errorf("want %s got %s", want, got)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		if err != ErrNoFunds {
			t.Errorf("want %s but got another", ErrNoFunds)
		}
	})
}
