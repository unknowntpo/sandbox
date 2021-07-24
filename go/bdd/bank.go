package main

import "errors"

type BankAccount struct {
	Balance int
}

func NewBankAccount(startingBalance int) *BankAccount {
	return &BankAccount{startingBalance}
}

func (account *BankAccount) Transfer(account2 *BankAccount, amount int) error {
	if account.Balance-amount < 0 {
		return errors.New("balance too low")
	}

	account.Balance -= amount
	account2.Balance += amount

	return nil
}
