package main

import "fmt"

func pointers2() {
	name := "Anirudh"
	var str *string = &name
	fmt.Println(*str)
}
