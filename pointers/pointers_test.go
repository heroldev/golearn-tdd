package pointers

import (
	"testing"
)

func TestWallet(test *testing.T) {

	test.Run("Deposit", func(t *testing.T) {

		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)

	})

	test.Run("Withdraw", func(t *testing.T) {

		wallet := Wallet{balance: Bitcoin(10)}

		err := wallet.Withdraw(Bitcoin(5))

		want := Bitcoin(5)

		assertBalance(t, wallet, want)
		assertNoError(t, err)

	})

	test.Run("Withdraw insufficient funds", func(t *testing.T) {

		startingBalance := Bitcoin(30)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(40))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficientFunds)

	})

}

func assertBalance(t testing.TB, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()
	if got != want {
		t.Errorf("got %q but wanted %q", got, want)
	}

}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got == nil {
		t.Errorf("wanted error but didn't get one")
	}

	if got != want {
		t.Errorf("got %q but want %q", got, want)
	}

}

func assertNoError(t testing.TB, got error) {
	t.Helper()

	if got != nil {
		t.Fatal("didn't want error but got one")
	}
}
