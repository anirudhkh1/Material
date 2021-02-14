package main

import "fmt"

func ranges() {
	//Iterating through slice using Range
	slice := []int{1, 2, 33, 4, 5, 5, 4, 45, 1}
	// fmt.Printf("Slice = %v", slice)

	// for _, e := range slice {
	// 	fmt.Printf("\nElement Value = %d", e)
	// }
	// fmt.Printf("\n")
	l := len(slice)
	for i, e := range slice[:l] {
		for _, e1 := range slice[i+1:] {
			if e == e1 {
				fmt.Printf("\n%v", e1)
			}
		}
	}

}
