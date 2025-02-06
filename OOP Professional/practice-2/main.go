package main

import (
	"fmt"
)

func main() {
	// // Initializing two bank accounts directly from the first file (bankAccount.go)
	// account1 := BankAccount{
	// 	AccountName:   "Nozibul Islam",
	// 	AccountNumber: 32300045,
	// 	Balance:       60000,
	// }

	// account2 := BankAccount{
	// 	AccountName:   "Ahmed Hossain",
	// 	AccountNumber: 54321098,
	// 	Balance:       20000,
	// }

	// // Displaying initial balances
	// fmt.Println("Initial Balance of account1:", account1.Balance)
	// fmt.Println("Initial Balance of account2:", account2.Balance)

	// // Deposit money into account1
	// account1.Deposit(5000)
	// fmt.Println("Account1 Balance after Deposit:", account1.Balance)

	// // Withdraw money from account1
	// success := account1.Withdraw(3000)
	// if success {
	// 	fmt.Println("Account1 Balance after Withdrawal:", account1.Balance)
	// } else {
	// 	fmt.Println("Account1 Withdrawal Failed (Insufficient Balance)")
	// }

	// Transfer money from account1 to account2
	// transferSuccess := account1.Transfer(4000, &account2)
	// if transferSuccess {
	// 	fmt.Println("Account1 Balance after Transfer:", account1.Balance)
	// 	fmt.Println("Account2 Balance after Transfer:", account2.Balance)
	// } else {
	// 	fmt.Println("Transfer Failed (Insufficient Balance in Account1)")
	// }

	// Now working with the second part of the account system (bankAccountPart2.go)
	account3 := BankAccountPart2{
		AccountName:   "John Doe",
		AccountNumber: 12345678,
		Balance:       10000,
	}

	account4 := BankAccountPart2{
		AccountName:   "Jane Doe",
		AccountNumber: 87654321,
		Balance:       15000,
	}

	// Deposit into account3
	account3.Deposit2(22000)
	fmt.Println("Account3 Balance after Deposit:", account3.Balance)

	// Withdraw from account3
	success2 := account3.Withdraw2(5000)
	if success2 {
		fmt.Println("Account3 Balance after Withdrawal:", account3.Balance)
	} else {
		fmt.Println("Account3 Withdrawal Failed (Insufficient Balance)")
	}

	// Transfer from account3 to account4
	transferSuccess2 := account3.Transfer2(3000, &account4)
	if transferSuccess2 {
		fmt.Println("Account3 Balance after Transfer:", account3.Balance)
		fmt.Println("Account4 Balance after Transfer:", account4.Balance)
	} else {
		fmt.Println("Transfer Failed (Insufficient Balance in Account3)")
	}
}
