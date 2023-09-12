package wallet

type Wallet struct {
	// lowercase means it is private
	// it is private outside the package it's defined in.
	balance int
}

// In Go, when you call a function or a method the arguments are copied.
// so if we pass w Wallet as arg, it is just a copy not the acutal one

// this means a pointer to wallet
func (w *Wallet) Deposit(amount int) {
	w.balance += amount
}

func (w *Wallet) Balance() int {
	return w.balance
	// we can also write (*w).balance to dereference
	// but GO permits us to write w.balance, without an explicit dereference.
	// These pointers to structs even have their own name: struct pointers
	// and they are automatically dereferenced
}
