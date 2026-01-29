package pointersanderrors

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10)) // Bitcoin型にキャスト

		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw with funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10)) // Bitcoin型にキャスト

		assertBalance(t, wallet, Bitcoin(10))
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		initialBalance := Bitcoin(20)
		wallet := Wallet{balance: initialBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, initialBalance)
		assertError(t, err, ErrInsufficientFunds)
	})
}

	func assertBalance(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()

		got := wallet.Balance()
		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	}

	func assertError(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("didn't get an error, but wanted one")
		}

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	func assertNoError(t *testing.T, got error) {
		t.Helper()

		if got != nil {
			t.Fatal("got an error, but didn't want one")
		}
	}
