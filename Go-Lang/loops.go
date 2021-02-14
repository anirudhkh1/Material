package main

import "fmt"

func loops() {
	x := 5
	for x < 8 {
		fmt.Printf("X = %d\n", x)
		x = x + 1
	}
	fmt.Println()
	var y int
	for y = 0; y <= 50; y++ {
		if y == 6 {
			continue
		}
		fmt.Printf("Y = %d\n", y)
	}
}
