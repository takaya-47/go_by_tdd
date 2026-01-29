package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10)) // Bitcoin型にキャスト

		got := wallet.Balance()
		want := Bitcoin(10)

		if got  != want {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		wallet.Withdraw(Bitcoin(10)) // Bitcoin型にキャスト

		got := wallet.Balance()
		want := Bitcoin(10)

		if got  != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}