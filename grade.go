package main


import "fmt"

//every course has different grading schemes, Core course → percentage (e.g., 83%)

//Lab course → pass/fail (e.g., PASS or FAIL)


// so we define a general interface that will be implemented by different grading types
// behaviour only -> pure polymorphism
type Grader interface{
	Grade(e Enrollment) (string,error)
}

type PercentageGrader struct{}

func (p *PercentageGrader) Grade(e Enrollment) (string,error){
	return fmt.Sprintf("%.1f%%",e.score*100),nil
	
}

type PassFailGrader struct{
	passMark float64 
}

func (p *PassFailGrader) Grade(e Enrollment) (string,error){
	if e.score>p.passMark{
		return "PASS" , nil 
	}
	return "FAIL" , nil 
}





