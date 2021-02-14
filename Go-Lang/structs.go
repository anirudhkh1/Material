package main

import "fmt"

type xyz struct {
	name   string
	age    int
	gender string
}

func changeName(abc *xyz) {
	abc.name = "GG BOIZ"
}

func structs() {
	// var Anirudh xyz = xyz{"Anirudh",18,"Male"}
	test := xyz{"Anirudh", 18, "Male"}
	fmt.Println(test.name, test.age, test.gender)
	test.name = "Vandana"
	fmt.Println(test.name, test.age, test.gender)
	test2 := xyz{gender: "Muskan", age: 22, name: "Female"}
	fmt.Println(test2.name, test2.age, test2.gender)

	//Changing Value through function
	test3 := xyz{"Sanjay", 55, "Male"}
	fmt.Println(test3)
	changeName(&test3)
	fmt.Println(test3)
}
