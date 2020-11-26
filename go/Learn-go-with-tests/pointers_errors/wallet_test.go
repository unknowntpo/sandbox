package wallet

import "testing"

func assertBalance(t *testing.T, got, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t *testing.T, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("Error expected, but got nothing")
	}

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	} else {
		t.Log("Error message is handled correctly")
	}
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("Expecting no error, but got one")
	}
}
func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}
		wallet.Deposit(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(20)

		assertBalance(t, got, want)
	})

	t.Run("Withdraw with sufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(10))
		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{balance: startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
		// Intended fail, check if error is triggered correctly
		assertError(t, err, ErrInsufficientFunds)
	})

}
