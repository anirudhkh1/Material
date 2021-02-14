package main

//Import Multiple Packages
import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func userInput() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Printf("Enter your Name = ")
	scanner.Scan()
	input := scanner.Text()
	fmt.Printf("Scanner Type = %T and Input type = %T\n", scanner, input)
	fmt.Printf("You typed = %q and type = %T", input, input)

	//Program to calculate Age
	fmt.Printf("\n\nEnter the year you were born in = ")
	scanner.Scan()
	input1, err := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("\nYou entered %d and your age is %d and Age type = %T.\n", input1, 2021-input1, input1)
	fmt.Printf("%v", err) //Prints out the Error

}
