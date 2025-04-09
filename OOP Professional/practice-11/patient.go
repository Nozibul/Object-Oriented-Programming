package main

import (
	"fmt"
	"time"
)

// Patient represents a hospital patient
type Patient struct {
	PatientID        int
	Name             string
	AdmissionDate    time.Time
	MedicalCondition string
	Ward             *Ward // Reference to the ward
}

// NewPatient creates a new patient
func NewPatient(id int, name string, medicalCondition string) *Patient {
	return &Patient{
		PatientID:        id,
		Name:             name,
		AdmissionDate:    time.Now(),
		MedicalCondition: medicalCondition,
		Ward:             nil,
	}
}

// GetDetails returns formatted patient information
func (p *Patient) GetDetails() string {
	wardInfo := "Not assigned"
	if p.Ward != nil {
		wardInfo = fmt.Sprintf("%s (ID: %d)", p.Ward.WardName, p.Ward.WardID)
	}

	return fmt.Sprintf("Patient ID: %d\nName: %s\nAdmission Date: %s\nMedical Condition: %s\nWard: %s",
		p.PatientID,
		p.Name,
		p.AdmissionDate.Format("2006-01-02 15:04"),
		p.MedicalCondition,
		wardInfo)
}

// TransferToWard transfers the patient to a different ward
func (p *Patient) TransferToWard(ward *Ward) error {
	if p.Ward == ward {
		return fmt.Errorf("patient is already in this ward")
	}

	return ward.AddPatient(p)
}

// DisplayPatientInfo prints information about the patient
func (p *Patient) DisplayPatientInfo() {
	fmt.Println("=== Patient Information ===")
	fmt.Println(p.GetDetails())
	fmt.Println("==========================")
}
