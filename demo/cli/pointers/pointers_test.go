package pointers

import "testing"

func TestDWallet(t *testing.T) {

	wallet := Wallet{}

	wallet.Deposit(Gencurrency(10))
	got := wallet.Balance()

	want := Gencurrency(10)

	if got != want {
		t.Errorf("got=%s want=%s", got, want)
	}
}

func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Gencurrency) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got=%s, expected=%s", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		Wallet := Wallet{}
		Wallet.Deposit(Gencurrency(10))
		assertBalance(t, Wallet, Gencurrency(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 10}
		wallet.Withdraw(Gencurrency(10))
		assertBalance(t, wallet, Gencurrency(0))
	})
}
