package main

import (
	"errors"
	"time"
)

// CreditCard represents a customer's credit card
type CreditCard struct {
	cardNumber     string
	expirationDate time.Time
	creditLimit    float64
	usedCredit     float64
	customerID     int // Reference back to the customer
}

// NewCreditCard creates a new credit card (constructor)
func NewCreditCard(cardNumber string, expirationDate time.Time, creditLimit float64, customerID int) (*CreditCard, error) {
	// Validate card number (must contain only digits)
	for _, char := range cardNumber {
		if char < '0' || char > '9' {
			return nil, errors.New("card number must contain only digits")
		}
	}

	// Validate credit limit
	if creditLimit <= 0 {
		return nil, errors.New("credit limit must be positive")
	}

	// Validate expiration date
	if expirationDate.Before(time.Now()) {
		return nil, errors.New("card is already expired")
	}

	return &CreditCard{
		cardNumber:     cardNumber,
		expirationDate: expirationDate,
		creditLimit:    creditLimit,
		usedCredit:     0.0,
		customerID:     customerID,
	}, nil
}

// Getters
func (c *CreditCard) GetCardNumber() string {
	return c.cardNumber
}

func (c *CreditCard) GetExpirationDate() time.Time {
	return c.expirationDate
}

func (c *CreditCard) GetCreditLimit() float64 {
	return c.creditLimit
}

func (c *CreditCard) GetUsedCredit() float64 {
	return c.usedCredit
}

func (c *CreditCard) GetCustomerID() int {
	return c.customerID
}

// IsExpired checks if the credit card is expired
func (c *CreditCard) IsExpired() bool {
	now := time.Now()
	return now.After(c.expirationDate)
}

// GetAvailableLimit returns the available credit limit
func (c *CreditCard) GetAvailableLimit() float64 {
	return c.creditLimit - c.usedCredit
}

// AddCharge adds a new charge to the card
func (c *CreditCard) AddCharge(amount float64) bool {
	// Check if there's enough available credit
	if c.usedCredit+amount > c.creditLimit {
		return false
	}

	c.usedCredit += amount
	return true
}

// MakePayment reduces the outstanding balance
func (c *CreditCard) MakePayment(amount float64) {
	if amount > c.usedCredit {
		c.usedCredit = 0
	} else {
		c.usedCredit -= amount
	}
}
