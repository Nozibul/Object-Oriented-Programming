package bank

import (
	"errors"
	"strconv"
)

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

// Constructor function for BankAccount
func NewBankAccount(name string, number int, balance int) (*BankAccount, error) {
	if name == "" {
		return nil, errors.New("account name cannot be empty")
	}
	if len(name) < 3 {
		return nil, errors.New("account name must be at least 3 characters long")
	}
	if len(strconv.Itoa(number)) < 6 {
		return nil, errors.New("account number must be at least 6 digits")
	}
	if balance < 0 {
		return nil, errors.New("initial balance cannot be negative")
	}

	return &BankAccount{
		accountName:   name,
		accountNumber: number,
		balance:       balance,
	}, nil
}

// Getter methods
// Rest of the methods remain the same but work with private fields
func (name *BankAccount) GetAccountName() string {
	return name.accountName
}
func (balance *BankAccount) GetBalance() int {
	return balance.balance
}

// Setter methods
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

// Deposit method
func (account *BankAccount) Deposit(amount int) {
	account.balance += amount
}

// Withdraw method
func (account *BankAccount) Withdraw(amount int) bool {
	if account.balance >= amount {
		account.balance -= amount
		return true
	}
	return false
}

// Transfer method
func (account *BankAccount) Transfer(amount int, targetAccount *BankAccount) bool {
	if account.Withdraw(amount) {
		targetAccount.Deposit(amount)
		return true
	}
	return false
}
