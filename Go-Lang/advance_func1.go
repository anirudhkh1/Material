package main

import "fmt"

func test1() func(string) {
	return func(b string) { fmt.Printf("\n%v", b) }
}

func advanceFunc1() {
	test1()("OHOH")
}
