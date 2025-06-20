package main

import "fmt"

type Student struct {
	name  string
	regno string
}

// Function to create a new student

func CreateNewStudent(name, regno string) Student {
	return Student{name: name, regno: regno}

}

// Creating getters for name and regno

func (s *Student) Name() string {
	return s.name
}

func (s *Student) Regno() string {
	return s.regno
}

func (s *Student) Display() {
	fmt.Printf("Student Regno: %v , name: %v", s.regno, s.name)
}
