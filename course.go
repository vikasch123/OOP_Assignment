package main

//import "fmt"



type Courses struct{
	id 	 int
	name string
}

func NewCourse(id int , name string) Courses{
	return Courses{id:id,name:name}
}


