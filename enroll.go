package main

import "fmt"

type Enrollment struct{
	Student 
	Courses
	Grader
	score float64
}


func NewEnrollment(st Student, c Courses, gr Grader,  score float64) Enrollment{
		return Enrollment{
		Student: st,
        Courses:  c,
        Grader:  gr,
        score:   score,
	}
}

func (e Enrollment) String() string{
	grade,_:= e.Grade(e)
	return fmt.Sprintf("%s ➜ %s : %s", e.Student.Name(), e.Courses.name, grade)
}

