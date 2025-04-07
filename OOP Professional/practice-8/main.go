package main

import (
	"fmt"
	"time"
)

func main() {
	// Example usage of the system
	fmt.Println("Bank Credit Card Management System")
	fmt.Println("----------------------------------")

	// Create a new customer using constructor
	birthDate := time.Date(1990, time.January, 15, 0, 0, 0, 0, time.UTC)
	customer, err := NewCustomer(1, "John Doe", birthDate)
	if err != nil {
		fmt.Println("Error creating customer:", err)
		return
	}
	fmt.Printf("Customer created: %s (ID: %d, Age: %d)\n",
		customer.GetName(), customer.GetID(), customer.GetAge())

	// Assign a credit card to the customer
	expirationDate := time.Date(2027, time.December, 31, 0, 0, 0, 0, time.UTC)
	err = customer.AssignCreditCard("1234567890123456", expirationDate, 500000)
	if err != nil {
		fmt.Println("Error assigning credit card:", err)
		return
	}
	fmt.Printf("Credit card assigned: %s (Limit: %.2f, Expires: %s)\n",
		customer.GetCreditCard().GetCardNumber(),
		customer.GetCreditCard().GetCreditLimit(),
		customer.GetCreditCard().GetExpirationDate().Format("02-01-2006"))

	// Make some purchases
	purchases := []float64{50000, 100000, 75000}
	for i, amount := range purchases {
		err = customer.MakePurchase(amount)
		if err != nil {
			fmt.Printf("Purchase %d failed: %s\n", i+1, err)
		} else {
			fmt.Printf("Purchase %d successful: %.2f\n", i+1, amount)
		}

		// Show current status after each purchase
		availableCredit, _ := customer.GetAvailableCredit()
		outstandingBalance, _ := customer.GetOutstandingBalance()
		fmt.Printf("  Available credit: %.2f, Outstanding balance: %.2f\n",
			availableCredit, outstandingBalance)
	}

	// Try to exceed credit limit
	fmt.Println("\nAttempting to exceed credit limit...")
	err = customer.MakePurchase(300000)
	if err != nil {
		fmt.Println("Purchase failed as expected:", err)
	} else {
		fmt.Println("Purchase successful (this should not happen)")
	}

	// Make a payment
	paymentAmount := 100000.0
	fmt.Printf("\nMaking payment: %.2f\n", paymentAmount)
	customer.GetCreditCard().MakePayment(paymentAmount)

	// Check final status
	availableCredit, _ := customer.GetAvailableCredit()
	outstandingBalance, _ := customer.GetOutstandingBalance()
	fmt.Printf("Final status - Available credit: %.2f, Outstanding balance: %.2f\n",
		availableCredit, outstandingBalance)

	// Check if card is valid
	valid, _ := customer.IsCardValid()
	if valid {
		fmt.Println("Credit card is valid")
	} else {
		fmt.Println("Credit card is expired")
	}
}
