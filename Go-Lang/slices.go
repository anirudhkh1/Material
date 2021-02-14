package main

import "fmt"

func slices() {
	arr := []int{1, 2, 3, 4, 5}
	var slicedArray []int = arr[:]
	fmt.Printf("Sliced Array = %v and type = %T and capacity = %d\n", slicedArray, slicedArray, cap(slicedArray))
	slicedArray = arr[1:3]
	fmt.Printf("Sliced Array = %v and type = %T and capacity = %d\n", slicedArray, slicedArray, cap(slicedArray))
	slicedArray = arr[2:]
	fmt.Printf("Sliced Array = %v and type = %T and capacity = %d\n", slicedArray, slicedArray, cap(slicedArray))

	var slice1 []int = []int{1, 2, 3, 4, 5}
	fmt.Printf("\nSlice1 = %v and Length = %d and Capacity = %d", slice1, len(slice1), cap(slice1))
	newSlice := append(slice1, 100)
	fmt.Printf("\nNew Slice = %v", newSlice)

	//Using Make Function
	sliceMake := make([]bool, 5)
	fmt.Printf("\nsliceMake = %v and Type = %T", sliceMake, sliceMake)

}
