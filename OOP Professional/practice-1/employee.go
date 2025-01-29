package main

import (
	"encoding/csv" // For reading and writing CSV files
	"fmt"          // For printing formatted output
	"os"           // For file and system operations (e.g., opening, closing files)
	"strconv"      // For converting strings to other data types (e.g., string to int)
)

// Employee struct to represent an employee's name and salary
type Employee struct {
	Name   string
	Salary int
}

// Global HandleError function to print error and return from the function
func HandleError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
}

// Read CSV File reads and parses a CSV file
func ReadCSVFile(filename string) ([][]string, error) {
	// Open CSV file
	file, err := os.Open(filename)
	HandleError(err) // Use HandleError for error handling

	// Read CSV file
	reader := csv.NewReader(file)
	lines, err := reader.ReadAll()
	HandleError(err) // Use HandleError for error handling

	return lines, nil
}

// ConvertLineToEmployee converts a CSV line into an Employee
func ConvertLineToEmployee(line []string) (Employee, error) {
	// Convert salary from string to integer
	salary, err := strconv.Atoi(line[1])
	if err != nil {
		HandleError(fmt.Errorf("error converting salary: %v", err)) // Use HandleError for error handling
		return Employee{}, err
	}

	// Return the Employee struct
	return Employee{Name: line[0], Salary: salary}, nil
}

// CalculateSalaryStats calculates the maximum, minimum, and average salary
func CalculateSalaryStats(employees []Employee) (Employee, Employee, float64) {
	if len(employees) == 0 {
		HandleError(fmt.Errorf("no employees found")) // Use HandleError for error handling
		return Employee{}, Employee{}, 0.0
	}

	// Initialize max and min employee with the first one in the list
	maxEmployee := employees[0]
	minEmployee := employees[0]
	totalSalary := 0

	// Iterate over all employees to find max, min, and total salary
	for _, emp := range employees {
		if emp.Salary > maxEmployee.Salary {
			maxEmployee = emp
		}
		if emp.Salary < minEmployee.Salary {
			minEmployee = emp
		}
		totalSalary += emp.Salary
	}

	// Calculate the average salary
	avgSalary := float64(totalSalary) / float64(len(employees))

	return maxEmployee, minEmployee, avgSalary
}

func main() {
	// Read the CSV data
	lines, err := ReadCSVFile("salarysheet.csv")
	HandleError(err) // Use HandleError for error handling

	// Parse the CSV lines into Employee structs
	var employees []Employee
	for _, line := range lines {
		emp, err := ConvertLineToEmployee(line)
		if err != nil {
			continue
		}
		employees = append(employees, emp)
	}

	// Calculate salary statistics
	maxEmployee, minEmployee, avgSalary := CalculateSalaryStats(employees)

	// Print the results
	if len(employees) > 0 {
		fmt.Printf("Maximum Salary: %s - %d\n", maxEmployee.Name, maxEmployee.Salary)
		fmt.Printf("Minimum Salary: %s - %d\n", minEmployee.Name, minEmployee.Salary)
		fmt.Printf("Average Salary: %.2f\n", avgSalary)
	} else {
		HandleError(fmt.Errorf("no employee data found")) // Use HandleError for error handling
	}
}
