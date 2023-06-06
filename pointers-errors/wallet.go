package main

import (
	"errors"
	"fmt"
)

type Bitcoin int

// Stringer is defined as a part of the format package
// https://pkg.go.dev/fmt#Stringer
type Stringer interface {
	String() string
}

// this is changing how we will see the "Bitcoin" type get printed
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin // lower case because private
}

func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if amount > w.balance {
		return errors.New("cannot withdraw, insufficient funds")
	}
	w.balance -= amount
	return nil
}
