package main

import "fmt"

func changeValue(str *string) {
	*str = "Changed!!"
}

func changeValue2(str string) string {
	str = "Changed2!!"
	return str
}

func pointers() {
	x := 5
	y := &x
	fmt.Printf("x=%v and y=%v and type of y=%T", x, y, y)
	*y = 22
	fmt.Printf("\nx=%v and y=%v and type of y=%T", x, y, y)
	z := *y
	fmt.Printf("\nz=%v", z)
	z = 23
	fmt.Printf("\nx=%v, y=%v and z=%v", x, y, z)

	//Lets see for Arrays
	fmt.Println()
	var a []int = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for i := 0; i < len(a); i++ {
		fmt.Printf("\na[%d]=%v and Address=%v", i, a[i], &a[i])
	}
	fmt.Printf("\n\n")

	//Using Functions
	name := "Anirudh Khurana"
	fmt.Println(name)
	changeValue(&name)
	fmt.Println(name)
	name = changeValue2(name)
	fmt.Println(name)

}
