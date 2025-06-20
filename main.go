package main
import "fmt"
func main() {
// ---------------- L1 -----------------
st := CreateNewStudent("1", "Alice Johnson")
c := NewCourse(101, "Intro to Go")
st.Display()
// ---------------- L2 -----------------
en1 := NewEnrollment(st, c, &PercentageGrader{}, 0.83)
en2 := NewEnrollment(st, c, &PassFailGrader{passMark: 0.5}, 0.83)
fmt.Println(en1)
fmt.Println(en2) // demonstrates polymorphism
}
