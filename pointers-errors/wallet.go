package main

import "fmt"

// Stringer is defined as a part of the format package
// https://pkg.go.dev/fmt#Stringer
type Stringer interface {
	String() string
}

type Bitcoin int

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
	return 0 // w.balance
}
