package main

// BankAccount struct definition
type BankAccount struct {
	AccountName            string
	AccountNumber, Balance int
}

// Deposit method adds to the balance
func (account *BankAccount) Deposit(amount int) {
	account.Balance += amount
}

// Withdraw method subtracts from the balance
func (account *BankAccount) Withdraw(amount int) bool {
	if account.Balance >= amount {
		account.Balance -= amount
		return true
	}
	return false // Not enough balance
}

// Transfer method moves an amount from one account to another
func (account *BankAccount) Transfer(amount int, targetAccount *BankAccount) bool {
	if account.Withdraw(amount) {
		targetAccount.Deposit(amount)
		return true
	}
	return false // Transfer failed due to insufficient balance
}

// func main() {
// 	// Initializing two bank accounts directly
// 	account1 := BankAccount{
// 		AccountName:   "Nozibul Islam",
// 		AccountNumber: 32300045,
// 		Balance:       60000,
// 	}

// 	account2 := BankAccount{
// 		AccountName:   "Ahmed Hossain",
// 		AccountNumber: 54321098,
// 		Balance:       20000,
// 	}

// 	// Displaying initial balances
// 	fmt.Println("Initial Balance of account1:", account1.Balance) // 60000
// 	fmt.Println("Initial Balance of account2:", account2.Balance) // 20000

// 	// Deposit money into account1
// 	account1.Deposit(5000)
// 	fmt.Println("Account1 Balance after Deposit:", account1.Balance) // 65000

// 	// Withdraw money from account1
// 	success := account1.Withdraw(3000)
// 	if success {
// 		fmt.Println("Account1 Balance after Withdrawal:", account1.Balance) // 62000
// 	} else {
// 		fmt.Println("Account1 Withdrawal Failed (Insufficient Balance)")
// 	}

// 	// Transfer money from account1 to account2
// 	transferSuccess := account1.Transfer(4000, &account2)
// 	if transferSuccess {
// 		fmt.Println("Account1 Balance after Transfer:", account1.Balance) // 58000
// 		fmt.Println("Account2 Balance after Transfer:", account2.Balance) // 24000
// 	} else {
// 		fmt.Println("Transfer Failed (Insufficient Balance in Account1)")
// 	}
// }
