package wallet

import (
	"errors"
	"fmt"
	"testing"
)

// 在 Go 中，错误是值，因此我们可以将其重构为一个变量，并为其提供一个单一的事实来源。
var InsufficientFundsError = errors.New("cannot withdraw, insufficient funds")

func assertNoError(t *testing.T, got error) {
	if got != nil {
		t.Fatal("got an error but didnt want one")
	}
}
func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want error) {
		if got == nil {
			t.Fatal("didn't get an error but wanted one")
		}
		if got != want {
			t.Errorf("got '%s', want '%s'", got, want)
		}
	}
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(50)}
		err := wallet.Withdraw(Bitcoin(60))
		assertBalance(t, wallet, Bitcoin(0))
		assertError(t, err, InsufficientFundsError)
	})
}

type Wallet struct {
	balance Bitcoin
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Println("address of balance in Deposit is", &w.balance)
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return InsufficientFundsError
	}
	w.balance -= amount
	return nil
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Bitcoin int
