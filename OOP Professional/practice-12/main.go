package main

import (
	"fmt"
	bank "mybank/inheritance"
)

func main() {
	// create savings account

	saving := bank.SavingsAccount{
		Account: bank.Account{
			AccountNo:    "SAV123555555",
			CustomerName: "Rahim",
			Address:      "Dhaka",
			Balance:      5000,
		},
		InterestRate:  .05,
		WithdrawLimit: 3,
	}

	// Create a chicking account
	checking := bank.CheckingAccount{
		Account: bank.Account{
			AccountNo:    "CHK456",
			CustomerName: "Karim",
			Address:      "Chittagong",
			Balance:      10000,
		},
		ServiceCharge: 100,
	}

	// Use savings account
	saving.Deposit(2000)
	saving.ApplyInterest()
	fmt.Printf("Savings Account Balance (Customer: %s): %.2f\n", saving.CustomerName, saving.Balance)

	// Use checking account
	checking.Withdraw(3000)
	checking.ApplyServiceCharge()
	fmt.Printf("Checking Account Balance (Customer: %s): %.2f\n", checking.CustomerName, checking.Balance)

}
