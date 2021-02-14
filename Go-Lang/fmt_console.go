package main

import "fmt"

func console() {
	var n1 int = 12
	fmt.Printf("\nNumber is %v and type is %T.", n1, n1)

	//Boolean
	var bo bool = true
	fmt.Printf("\n\nThis is a boolean value = %t.", bo)

	//Converting to Base i.e 2(%b),8(%o),10(%d),16(%x)
	fmt.Printf("\n\nConverting %v to Binary = %b, Octal = %o, Decimal = %d and Hexadecimal = %x.", n1, n1, n1, n1, n1)

	//Floating Point (%e,%f,%F,%g)
	var n2 float64 = 45.4534423432432
	fmt.Printf("\n\nNumber is %f and type is %T.", n2, n2)

	//Strings
	fmt.Printf("\nString = %s and String in double quotes = %q", "Anirudh", "Anirudh")

	//Width and Precision
	var name string = "Ani"
	var x float32 = 45.65467765
	fmt.Printf("\nBelow Name with 9 Padding\n%9.2s\n\n", name)
	fmt.Printf("\nX = %.2f", x)
	fmt.Printf("\nX = %015.2f", x)
}
