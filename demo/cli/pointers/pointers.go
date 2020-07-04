package pointers

import "fmt"

// Gencurrency irepresents number of Gencurrency
type Gencurrency int

func (g Gencurrency) String() string {
	return fmt.Sprintf("%d Gcurr", g)
}

// Wallet stores Gencurrency
type Wallet struct {
	balance Gencurrency
}

// Deposit will add Gencurrency
func (w *Wallet) Deposit(amount Gencurrency) {
	w.balance += amount
}

// Withdraw will deduct amount
func (w *Wallet) Withdraw(amount Gencurrency) {
	w.balance -= amount
}

//Balance returns amount in wallet
func (w *Wallet) Balance() Gencurrency {
	return w.balance
}
