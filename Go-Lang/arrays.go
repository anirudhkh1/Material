package main

import "fmt"

func arrays() {
	var arr [5]bool
	arr1 := []bool{true}
	arr2 := [5]bool{}
	arr4 := [3]int{1, 2, 3}
	fmt.Printf("Array = %v and Type = %T\n", arr, arr)
	fmt.Printf("Array = %v and Type = %T\n", arr1, arr1)
	fmt.Printf("Array = %v and Type = %T\n", arr2, arr2)
	fmt.Printf("Array = %v and Type = %T\n\n", arr4, arr4)

	arr4[0] = 100
	fmt.Printf("Array 4 = %v and lenget of Array = %d", arr4, len(arr4))

	//Sum of all elements
	sum := 0
	for i := 0; i < len(arr4); i++ {
		sum += arr4[i]
	}
	fmt.Printf("\nTotal sum of Array %v = %d", arr4, sum)

	//2D Array
	arr5 := [3][3]int{{1, 2}, {3, 4}}
	fmt.Printf("\nArray 5 = %v and type = %T", arr5, arr5)

	var arr6 [5]int = [5]int{1, 2, 33, 34, 45}
	fmt.Printf("\nArray = %v", arr6)
}
