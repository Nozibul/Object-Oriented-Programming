package bank

import "strconv"

type ErrorStruct struct {
	message string
}

func (e *ErrorStruct) Error() string {
	return e.message
}

// BankAccount struct with private fields (lowercase)
type BankAccount struct {
	accountName   string // private field
	accountNumber int    // private field
	balance       int    // private field
}

// Rest of the methods remain the same but work with private fields
func (name *BankAccount) GetAccountName() string {
	return name.accountName
}

func (name *BankAccount) SetAccountName(newName string) error {
	if newName == "" {
		return &ErrorStruct{"Account name cannot be empty"}
	}
	if len(newName) < 3 {
		return &ErrorStruct{"Account name must be at least 3 characters long"}
	}
	name.accountName = newName
	return nil
}

func (number *BankAccount) SetAccountNumber(newNumber int) error {
	if len(strconv.Itoa(newNumber)) < 6 {
		return &ErrorStruct{"Account number must be at least 6 digits"}
	}
	number.accountNumber = newNumber
	return nil
}

func (balance *BankAccount) GetBalance() int {
	return balance.balance
}

func (account *BankAccount) Deposit(amount int) {
	account.balance += amount
}

func (account *BankAccount) Withdraw(amount int) bool {
	if account.balance >= amount {
		account.balance -= amount
		return true
	}
	return false
}

func (account *BankAccount) Transfer(amount int, targetAccount *BankAccount) bool {
	if account.Withdraw(amount) {
		targetAccount.Deposit(amount)
		return true
	}
	return false
}
