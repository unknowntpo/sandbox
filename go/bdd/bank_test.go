package main

import (
	"fmt"
	"testing"
)

func TestBankAccount(t *testing.T) {
	var account1, account2 *BankAccount
	var err error

	AccountHasBalance := func(account **BankAccount, balance int) stepFunc {
		return stepFunc(func(t *testing.T) {
			*account = NewBankAccount(balance)
		})
	}

	ITransfer20DollarsBetweenAccounts := stepFunc(func(t *testing.T) {
		err = account1.Transfer(account2, 20)
	})

	AccountShouldHaveBalance := func(account **BankAccount, balance int) stepFunc {
		return func(t *testing.T) {
			if got := (*account).Balance; got != balance {
				t.Errorf("got %d, wanted %d", got, balance)
			}
		}
	}

	IShouldReceiveAnError := func(expected string) stepFunc {
		return stepFunc(func(t *testing.T) {
			if got := fmt.Sprintf("%v", err); got != expected {
				t.Errorf("got %v, wanted %s", got, expected)
			}
		})
	}

	Scenario(t, "Transfering between accounts",
		Given(AccountHasBalance(&account1, 100)),
		And(AccountHasBalance(&account2, 40)),
		When(ITransfer20DollarsBetweenAccounts),
		Then(AccountShouldHaveBalance(&account1, 80)),
		And(AccountShouldHaveBalance(&account2, 60)))

	Scenario(t, "Balance too low to transfer",
		Given(AccountHasBalance(&account1, 10)),
		And(AccountHasBalance(&account2, 40)),
		When(ITransfer20DollarsBetweenAccounts),
		Then(IShouldReceiveAnError("balance too low")),
		And(AccountShouldHaveBalance(&account1, 10)),
		And(AccountShouldHaveBalance(&account2, 40)))
}

// ---
// Everything below this line is just to run the tests.

func Scenario(t *testing.T, name string, steps ...stepFunc) {
	t.Run(name, func(t *testing.T) {
		for _, step := range steps {
			step(t)
		}
	})
}

type stepFunc func(t *testing.T)

func step(fn stepFunc) stepFunc {
	return fn
}

var (
	Given = step
	When  = step
	Then  = step
	And   = step
)
