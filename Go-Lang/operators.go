package main

import "fmt"

func operators() {
	var x int = 10
	var y int = 6
	fmt.Printf("Addition = %d\nSub = %d\nDiv = %f", x+y, x-y, float64(x)/float64(y))
}
