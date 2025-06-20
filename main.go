package main

import "fmt"

func main() {
	// ---------------- L1 -----------------
	st := CreateNewStudent("Alice", "1")
	c1 := NewCourse(101, "Intro to Go")
	c2 := NewCourse(102, "Networking Lab")

	reg := &Registrar{}
	reg.AddStudent(st)
	reg.AddCourse(c1)
	reg.AddCourse(c2)

	reg.Enroll(NewEnrollment(st, c1, &PercentageGrader{}, 0.91))
	reg.Enroll(NewEnrollment(st, c2, &PassFailGrader{0.7}, 0.65))

	// ---------------- L2 -----------------
	en1 := NewEnrollment(st, c1, &PercentageGrader{}, 0.83)
	en2 := NewEnrollment(st, c2, &PassFailGrader{passMark: 0.5}, 0.83)
	fmt.Println(en1)
	fmt.Println(en2) // demonstrates polymorphism

	reg.SetGrader(c2.id, &PercentageGrader{})

	for _, e := range reg.Enrollments() {
		fmt.Println(e)
	}

	if err := ExportTranscript("transcript.csv", reg.Enrollments()); err != nil {
		fmt.Println("csv error:", err)
	}

}
