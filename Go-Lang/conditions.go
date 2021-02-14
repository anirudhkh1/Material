package main

import "fmt"

func conditions() {
	fmt.Printf("%t", !!(true || false))
	fmt.Printf("\n%t", false && true)
}
