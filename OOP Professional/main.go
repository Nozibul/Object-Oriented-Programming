package main

import (
	"bank-project/bank"
	"fmt"
)

func main() {
	// Create a new bank account
	account3, err := bank.NewBankAccount("Nozibul", 552225, 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Deposit money
	account3.Deposit(20000)

	// Set account name with error checking
	err = account3.SetAccountName("ABsssss")
	if err != nil {
		fmt.Println("Error setting name:", err)
		return
	}

	// Print account information
	fmt.Printf("Account Name: %s\n", account3.GetAccountName())
	fmt.Printf("Balance: %d\n", account3.GetBalance())

	// Create second account for transfer demo
	account4 := &bank.BankAccount{}
	account4.SetAccountName("Jane Doe")
	account4.SetAccountNumber(87654321)

	// Transfer money between accounts
	if account3.Transfer(10000, account4) {
		fmt.Printf("Account3 Balance: %d\n", account3.GetBalance())
		fmt.Printf("Account4 Balance: %d\n", account4.GetBalance())
	} else {
		fmt.Println("Transfer failed - insufficient balance")
	}
}
