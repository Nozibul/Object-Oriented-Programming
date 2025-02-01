package main

import (
	"fmt"
)

// Error interface definition
type Error interface {
	Error() string
}

// Error struct for custom error messages
type ErrorStruct struct {
	Message string
}

// Implementing the Error method for ErrorStruct
func (e *ErrorStruct) Error() string {
	return e.Message
}

// Constants for limits
const (
	MaxLimit         = 500000 // Max Limit for spending
	DailyLimit       = 100000 // Daily Withdraw Limit
	TransactionLimit = 20000  // Per Transaction Limit
)

// CreditCard struct definition
type CreditCard struct {
	Name           string
	Number         int
	Balance        int
	TotalSpent     int
	DailyWithdrawn int
}

// Withdraw method (টাকা উত্তোলনের জন্য)
func (c *CreditCard) Withdraw(amount int) Error {
	// Checking if the withdrawal exceeds the per-transaction limit
	if amount > TransactionLimit {
		return &ErrorStruct{"Transaction exceeds per-transaction limit of 20,000"}
	}

	// Checking if the daily withdrawal limit is exceeded
	if c.DailyWithdrawn+amount > DailyLimit {
		return &ErrorStruct{"Daily withdrawal limit exceeded (Max 100,000 per day)"}
	}

	// Checking if the balance is enough for withdrawal
	if amount > c.Balance {
		return &ErrorStruct{"Insufficient balance"}
	}

	// Perform withdrawal and update daily limit and balance
	c.Balance -= amount
	c.DailyWithdrawn += amount
	c.TotalSpent += amount

	fmt.Println("✅ Withdrawn:", amount)
	fmt.Println("💰 Remaining Balance:", c.Balance)
	return nil
}

// Bill Payment method (বিল পেমেন্টের জন্য)
func (c *CreditCard) PayBill(amount int) Error {
	// Checking if total spending exceeds the max limit
	if c.TotalSpent+amount > MaxLimit {
		return &ErrorStruct{"Total spending exceeds the maximum limit of 500,000"}
	}

	// Perform bill payment and update total spending
	c.TotalSpent += amount
	fmt.Println("✅ Bill Paid:", amount)
	fmt.Println("💰 Remaining Balance:", c.Balance)
	return nil
}

func main() {
	// CreditCard struct-এর একটি instance তৈরি করা হলো
	card := CreditCard{
		Name:    "Nozibul Islam",
		Number:  12255789,
		Balance: 400000, // User gives balance
	}

	// Withdraw 20,000 (Valid Transaction)
	if err := card.Withdraw(20000); err != nil {
		fmt.Println(err.Error()) // Error handling
	} else {
		fmt.Println("💰 Remaining Balance:", card.Balance)
	}

	// Try another withdrawal of 80,000 (Valid Transaction)
	if err := card.Withdraw(80000); err != nil {
		fmt.Println(err.Error()) // Error handling
	} else {
		fmt.Println("💰 Remaining Balance:", card.Balance)
	}

	// Try to withdraw more than the daily limit (Invalid Transaction)
	if err := card.Withdraw(20000); err != nil {
		fmt.Println(err.Error()) // Error handling
	} else {
		fmt.Println("💰 Remaining Balance:", card.Balance)
	}

	// Pay Bill (Valid Bill Payment)
	if err := card.PayBill(50000); err != nil {
		fmt.Println(err.Error()) // Error handling
	} else {
		fmt.Println("💰 Remaining Balance:", card.Balance)
	}

	// Try to exceed the total spending limit (Invalid Bill Payment)
	if err := card.PayBill(400000); err != nil {
		fmt.Println(err.Error()) // Error handling
	} else {
		fmt.Println("💰 Remaining Balance:", card.Balance)
	}
}
