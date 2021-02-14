package main

import "fmt"

func mapss() {
	var mp map[string]int = map[string]int{
		"Ani":    18,
		"Muskan": 19,
	}
	fmt.Printf("MP = %v", mp)
	fmt.Printf("\n%v", mp["Ani"])
	mp["Ani"] = 222
	fmt.Printf("\n%v", mp)
	delete(mp, "Ani")
	fmt.Printf("\n%v", mp)
	mp["Vandana"] = 222
	fmt.Printf("\n%v", mp)

	//Check if value Exists
	val, exist := mp["Vandaana"]
	fmt.Printf("\nVal = %v and Exist = %v", val, exist)

}
