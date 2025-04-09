package main

import (
	"fmt"
)

func main() {
	// Create wards
	cardiology := NewWard(1, "Cardiology", "Dr. Rahman", 3)
	orthopedic := NewWard(2, "Orthopedic", "Dr. Akter", 2)

	// Create patients
	karim := NewPatient(101, "Abdul Karim", "Heart Attack")
	rahim := NewPatient(102, "Rahim Ahmed", "Chest Pain")
	salma := NewPatient(103, "Salma Begum", "Fractured Leg")
	fatima := NewPatient(104, "Fatima Akter", "Back Pain")
	jasim := NewPatient(105, "Jasim Uddin", "Cardiac Arrhythmia")

	// Add patients to wards
	fmt.Println("Adding patients to wards...")

	err := cardiology.AddPatient(karim)
	printOperationResult("Add Karim to Cardiology", err)

	err = cardiology.AddPatient(rahim)
	printOperationResult("Add Rahim to Cardiology", err)

	err = orthopedic.AddPatient(salma)
	printOperationResult("Add Salma to Orthopedic", err)

	err = orthopedic.AddPatient(fatima)
	printOperationResult("Add Fatima to Orthopedic", err)

	// Try to add a patient to a full ward
	err = orthopedic.AddPatient(jasim)
	printOperationResult("Add Jasim to Orthopedic (already full)", err)

	// Display ward information
	fmt.Println("\n=== Ward Status ===")
	cardiology.DisplayWardInfo()
	orthopedic.DisplayWardInfo()

	// Transfer a patient from one ward to another
	fmt.Println("\nTransferring Rahim from Cardiology to Orthopedic...")
	// First remove a patient from orthopedic to make space
	err = orthopedic.RemovePatient(salma.PatientID)
	printOperationResult("Remove Salma from Orthopedic", err)

	// Now try the transfer
	err = rahim.TransferToWard(orthopedic)
	printOperationResult("Transfer Rahim to Orthopedic", err)

	// Display updated ward information
	fmt.Println("\n=== Updated Ward Status ===")
	cardiology.DisplayWardInfo()
	orthopedic.DisplayWardInfo()

	// Display individual patient info
	fmt.Println("\n=== Patient Details ===")
	rahim.DisplayPatientInfo()
	salma.DisplayPatientInfo() // Salma has been removed from any ward
}

func printOperationResult(operation string, err error) {
	if err != nil {
		fmt.Printf("❌ %s: %s\n", operation, err)
	} else {
		fmt.Printf("✅ %s: Success\n", operation)
	}
}
