package main

import "fmt"

type student struct {
	name  string
	age   int
	grade []int
}

func (s student) getAge() int {
	return s.age
}

func (s *student) setAge(age int) {
	s.age = age
}

func (s student) averageGrades() float64 {
	sum := 0
	for _, v := range s.grade {
		sum += v
	}
	return float64(sum) / 3.0
}

func structmethods() {
	anirudh := student{"Anirudh", 33, []int{33, 55, 66}}

	fmt.Println(anirudh)
	fmt.Println(anirudh.getAge())
	anirudh.setAge(3)
	fmt.Println(anirudh)
	fmt.Printf("Average grades = %.2f", anirudh.averageGrades())
}
