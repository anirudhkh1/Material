package main

import "fmt"

func variables() {
	// Implicit
	var name string = "Anirudh"
	var number uint = 0
	fmt.Printf("\nType of number = %T and name = %T", number, name)
	fmt.Println()

	// Explicit
	var number2 = 2555
	fmt.Printf("\nNumber2 has value = %v and type = %T", number2, number2)
	// Another Way
	number3 := 224324
	fmt.Printf("\nNumber3 has value = %v and type = %T", number3, number3)

	// When we defice variable but do not give it a value
	var number4 uint16
	var charr string
	var boo bool

	fmt.Printf("\nValue of number4 = %v, charr = %v and boo = %v", number4, charr, boo)

}
