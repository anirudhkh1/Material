package main

import "fmt"

func test() {
	fmt.Printf("Hello World")
}

func advance(myFunc func(int) int) {
	fmt.Printf("\n\n%d", myFunc(6))
}

func advanceFunc() {
	defer fmt.Printf("\n\nDone with functions.")
	x := test
	x()

	y := func(a string) {
		fmt.Printf("\n%v", a)
	}
	y("Anirudh")

	stored := func(a int) int {
		return a
	}
	// fmt.Printf("\nStored = %d", stored)

	advance(stored)

	//Just call a function
	func() {
		fmt.Printf("\nJust a function")
	}()

}
