package main

import "fmt"

func abc() {
	ans1, ans2 := sum(2, 3)
	fmt.Printf("Addition = %d and Sub = %d", ans1, ans2)
	ans3, ans4 := mul(5, 7)
	fmt.Printf("\nAddition = %d and Sub = %d", ans3, ans4)
}

func sum(a, b int) (int, int) {
	return a + b, a - b
}

//Labelled Return

func mul(a, b int) (a1, a2 int) {
	defer fmt.Printf("\nLine1")
	a1 = a + b
	fmt.Printf("\nLine2")
	a2 = a - b
	fmt.Printf("\nLine3")
	return
}
