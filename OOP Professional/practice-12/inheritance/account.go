package inheritance

import "fmt"

type Account struct {
	AccountNo    string
	CustomerName string
	Address      string
	Balance      float64
}

// Deposite method adds to the balance
func (a *Account) Deposit(amount float64) {
	a.Balance += amount
}

// withdraw method subtracts from the balance
func (a *Account) Withdraw(amount float64) bool {
	if amount > a.Balance {
		fmt.Println("Insufficient balance")
	}
	a.Balance -= amount
	return true
}
