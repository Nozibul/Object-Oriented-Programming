package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

// Static variables
var (
	totalStudentsEnrolled int = 0
	PassingMark           int = 40
)

// Constructor function for Student (like a class constructor)
func NewStudent(name string, age int) *Student {
	totalStudentsEnrolled++
	return &Student{
		Name: name,
		Age:  age,
	}
}

// GetNoOfStudentsEnrolled returns the total number of students created
func GetNoOfStudentsEnrolled() int {
	return totalStudentsEnrolled
}

func main() {
	// Create some students
	student1 := NewStudent("Rahul", 20)
	student2 := NewStudent("Priya", 22)

	// Print student details
	fmt.Printf("Student 1: Name = %s, Age = %d\n", student1.Name, student1.Age)
	fmt.Printf("Student 2: Name = %s, Age = %d\n", student2.Name, student2.Age)

	// Print total number of students
	fmt.Printf("Total Students Enrolled: %d\n", GetNoOfStudentsEnrolled())

	// Print passing mark
	fmt.Printf("Passing Mark: %d\n", PassingMark)
}
