package main

import "testing"

func TestPassFailGrader(t * testing.T){

	e:=Enrollment{score:0.4,Grader:&PassFailGrader{0.6}}
	msg,_ :=e.Grade(e)
	if msg!="FAIL"{
		t.Fatal("Expected FAILED message")
	}

	e1:=Enrollment{score:0.7,Grader:&PassFailGrader{0.6}}


	msg1,_:=e1.Grade(e1)

	if msg1!="PASS"{
		t.Fatal("Expected PASS message")
	}


}

func TestPercentageGrader(t *testing.T){
	e:=Enrollment{score:0.875,Grader:&PercentageGrader{}}
	got,_:=e.Grade(e)
	want:="87.5%"

	if got!=want{
		t.Fatalf("got %s want %s",got,want)


	}

}
