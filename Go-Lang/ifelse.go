package main

import "fmt"

func ifelse() {
	name := 18
	if name > 18 {
		fmt.Println("You can drive.")
	} else if name < 18 {
		fmt.Println("You cannot drive.")
	} else if name == 18 {
		fmt.Println("You are just 18. Practice Driving.")
	}
}
