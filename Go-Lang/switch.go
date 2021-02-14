package main

import "fmt"

func switchh() {
	ans := 3
	switch ans {
	case 3:
		fmt.Println("Its 3")

	default:
		fmt.Println("Not a Case")
	}

	switch {
	case ans > 3:
		fmt.Println("Greater than 3.")
	case ans == 3:
		fmt.Println("Equal to 3.")
	}
}
