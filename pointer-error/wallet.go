package wallet

import (
	"errors"
	"fmt"
)

var ErrInsufficientFunds = errors.New("cannot withdraw, insufficient fund")

type Bitcoin int

type Stringer interface {
	String() string
}

type Wallet struct {
	// lowercase means it is private
	// it is private outside the package it's defined in.
	balance Bitcoin
}

// In Go, when you call a function or a method the arguments are copied.
// so if we pass w Wallet as arg, it is just a copy not the acutal one

// this means a pointer to wallet
func (w *Wallet) Deposit(amount Bitcoin) {
	w.balance += amount
}

func (w *Wallet) Withdraw(amount Bitcoin) error {
	if w.balance >= amount {
		w.balance -= amount
		return nil
	}
	// errors.New creates a new error with a message of your choosing.
	return ErrInsufficientFunds
}

func (w *Wallet) Balance() Bitcoin {
	return w.balance
	// we can also write (*w).balance to dereference
	// but GO permits us to write w.balance, without an explicit dereference.
	// These pointers to structs even have their own name: struct pointers
	// and they are automatically dereferenced
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
