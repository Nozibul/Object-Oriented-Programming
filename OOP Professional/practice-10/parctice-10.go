package main

import (
	"errors"
	"fmt"
	"time"
)

// Student represents a university student
type Student struct {
	StudentID     string
	Name          string
	AdmissionDate time.Time
	Department    *Department
}

// Department represents a university department
type Department struct {
	DepartmentName   string
	HeadOfDepartment string
	MaximumCapacity  int
	Students         []*Student
}

// NewStudent creates a new student
func NewStudent(id string, name string, admissionDate time.Time) *Student {
	return &Student{
		StudentID:     id,
		Name:          name,
		AdmissionDate: admissionDate,
	}
}

// NewDepartment creates a new department
func NewDepartment(name string, head string, capacity int) *Department {
	return &Department{
		DepartmentName:   name,
		HeadOfDepartment: head,
		MaximumCapacity:  capacity,
		Students:         make([]*Student, 0),
	}
}

// AddStudent adds a student to the department if there is capacity
func (d *Department) AddStudent(student *Student) error {
	// Check if department has capacity
	if len(d.Students) >= d.MaximumCapacity {
		return errors.New("department has reached maximum capacity")
	}

	// Check if student already belongs to this department
	for _, s := range d.Students {
		if s.StudentID == student.StudentID {
			return errors.New("student already belongs to this department")
		}
	}

	// Check if student already belongs to another department
	if student.Department != nil {
		return errors.New("student already belongs to another department")
	}

	// Add student to department
	d.Students = append(d.Students, student)
	student.Department = d

	return nil
}

// RemoveStudent removes a student from the department
func (d *Department) RemoveStudent(studentID string) error {
	for i, student := range d.Students {
		if student.StudentID == studentID {
			// Remove student from department
			d.Students = append(d.Students[:i], d.Students[i+1:]...)
			student.Department = nil
			return nil
		}
	}
	return errors.New("student not found in department")
}

// GenerateStudentList returns a list of all students in the department
func (d *Department) GenerateStudentList() []*Student {
	return d.Students
}

// GetCurrentStudentCount returns the current number of students in the department
func (d *Department) GetCurrentStudentCount() int {
	return len(d.Students)
}

func main() {
	// Create departments
	csDept := NewDepartment("Computer Science", "Dr. Smith", 3) // Small capacity for testing
	mathDept := NewDepartment("Mathematics", "Dr. Johnson", 2)

	// Create students
	student1 := NewStudent("S001", "Alice", time.Now())
	student2 := NewStudent("S002", "Bob", time.Now())
	student3 := NewStudent("S003", "Charlie", time.Now())
	student4 := NewStudent("S004", "Diana", time.Now())

	// Add students to departments
	fmt.Println("Adding students to departments:")

	err := csDept.AddStudent(student1)
	printResult("Add Alice to CS", err)

	err = csDept.AddStudent(student2)
	printResult("Add Bob to CS", err)

	err = mathDept.AddStudent(student3)
	printResult("Add Charlie to Math", err)

	// Try to add a student to a department that's already at capacity
	err = mathDept.AddStudent(student4)
	printResult("Add Diana to Math (should fail due to capacity)", err)

	// Try to add a student to multiple departments
	err = mathDept.AddStudent(student1)
	printResult("Add Alice to Math (should fail as already in CS)", err)

	// Print department info
	fmt.Println("\nDepartment Information:")
	printDepartmentInfo(csDept)
	printDepartmentInfo(mathDept)

	// Remove a student
	fmt.Println("\nRemoving Bob from CS department:")
	err = csDept.RemoveStudent("S002")
	printResult("Remove Bob from CS", err)

	// Print updated department info
	fmt.Println("\nUpdated Department Information:")
	printDepartmentInfo(csDept)

	// Now we can add Diana to Math after removing someone
	err = csDept.AddStudent(student4)
	printResult("Add Diana to CS after removing Bob", err)

	fmt.Println("\nFinal Department Information:")
	printDepartmentInfo(csDept)
	printDepartmentInfo(mathDept)
}

func printResult(action string, err error) {
	if err != nil {
		fmt.Printf("%s: FAILED - %s\n", action, err)
	} else {
		fmt.Printf("%s: SUCCESS\n", action)
	}
}

func printDepartmentInfo(dept *Department) {
	fmt.Printf("%s Department (Head: %s)\n", dept.DepartmentName, dept.HeadOfDepartment)
	fmt.Printf("- Capacity: %d/%d\n", dept.GetCurrentStudentCount(), dept.MaximumCapacity)
	fmt.Println("- Students:")

	for _, student := range dept.Students {
		fmt.Printf("  * %s: %s (Admitted: %s)\n",
			student.StudentID,
			student.Name,
			student.AdmissionDate.Format("2006-01-02"))
	}
	fmt.Println()
}
