package main

import (
	"encoding/csv"
	//"fmt"
	"os"
)

func ExportTranscript(path string, list []Enrollment) error {
	file, err := os.Create(path)
	if err != nil {
		panic(err)
	}

	defer file.Close()
	w := csv.NewWriter(file)
	defer w.Flush()

	_ = w.Write([]string{"Student", "Course", "Grade"})

	for _, e := range list {
		grade, _ := e.Grade(e)
		_ = w.Write([]string{e.Student.Name(), e.Courses.name, grade})

	}

	return w.Error()
}
