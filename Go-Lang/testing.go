package main

import "fmt"

func init() {
	fmt.Println("In Init")
}

func testing() {

	defer fmt.Println("Defer")

	fmt.Println("Hello World!!")
}
