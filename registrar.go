package main

type Registrar struct {
	students    []Student
	courses     []Courses
	enrollments []Enrollment
}

func (r *Registrar) AddStudent(s Student) {
	r.students = append(r.students, s)
}

func (r *Registrar) AddCourse(c Courses) {
	r.courses = append(r.courses, c)
}

func (r *Registrar) Enroll(e Enrollment) {
	r.enrollments = append(r.enrollments, e)
}

// setters for stratergy replacements

func (r *Registrar) SetGrader(id int, g Grader) {
	for i, e := range r.enrollments {
		if e.Courses.id == id {
			r.enrollments[i].Grader = g
		}
	}
}

func (r Registrar) Enrollments() []Enrollment {
	return r.enrollments
}
