package main

import (
	"errors"
	"fmt"
)

// Ward represents a hospital ward
type Ward struct {
	WardID      int
	WardName    string
	HeadNurse   string
	MaxCapacity int
	Patients    []*Patient
}

// NewWard creates a new ward
func NewWard(id int, name string, headNurse string, maxCapacity int) *Ward {
	return &Ward{
		WardID:      id,
		WardName:    name,
		HeadNurse:   headNurse,
		MaxCapacity: maxCapacity,
		Patients:    make([]*Patient, 0),
	}
}

// AddPatient adds a patient to the ward if capacity allows
func (w *Ward) AddPatient(patient *Patient) error {
	if w.IsAtCapacity() {
		return errors.New("cannot add patient: ward is at maximum capacity")
	}

	// Check if patient is already in this ward
	for _, p := range w.Patients {
		if p.PatientID == patient.PatientID {
			return errors.New("patient already exists in this ward")
		}
	}

	// Remove patient from previous ward if any
	if patient.Ward != nil {
		patient.Ward.RemovePatient(patient.PatientID)
	}

	// Add patient to this ward
	w.Patients = append(w.Patients, patient)
	patient.Ward = w
	return nil
}

// RemovePatient removes a patient from the ward
func (w *Ward) RemovePatient(patientID int) error {
	for i, patient := range w.Patients {
		if patient.PatientID == patientID {
			// Remove patient from slice
			w.Patients = append(w.Patients[:i], w.Patients[i+1:]...)
			patient.Ward = nil
			return nil
		}
	}
	return errors.New("patient not found in ward")
}

// IsAtCapacity checks if the ward is at maximum capacity
func (w *Ward) IsAtCapacity() bool {
	return len(w.Patients) >= w.MaxCapacity
}

// GetPatientList returns list of all patients in the ward
func (w *Ward) GetPatientList() []*Patient {
	return w.Patients
}

// DisplayWardInfo prints information about the ward
func (w *Ward) DisplayWardInfo() {
	fmt.Printf("Ward ID: %d\n", w.WardID)
	fmt.Printf("Ward Name: %s\n", w.WardName)
	fmt.Printf("Head Nurse: %s\n", w.HeadNurse)
	fmt.Printf("Capacity: %d/%d\n", len(w.Patients), w.MaxCapacity)
	fmt.Printf("=== Patients (%d) ===\n", len(w.Patients))

	for _, patient := range w.Patients {
		fmt.Printf("- %s (ID: %d)\n", patient.Name, patient.PatientID)
	}
	fmt.Println("==================")
}
