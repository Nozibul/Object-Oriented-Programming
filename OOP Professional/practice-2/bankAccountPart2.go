package main

// BankAccountPart2 struct definition
type BankAccountPart2 struct {
	AccountName            string
	AccountNumber, Balance int
}

// Deposit2 method adds to the balance
func (account *BankAccountPart2) Deposit2(amount int) {
	account.Balance += amount
}

// Withdraw2 method subtracts from the balance
func (account *BankAccountPart2) Withdraw2(amount int) bool {
	if account.Balance >= amount {
		account.Balance -= amount
		return true
	}
	return false
}

// Transfer2 method moves an amount from one account to another
func (account *BankAccountPart2) Transfer2(amount int, targetAccount *BankAccountPart2) bool {
	if account.Withdraw2(amount) {
		targetAccount.Deposit2(amount)
		return true
	}
	return false
}
