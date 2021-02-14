package main

import "fmt"

func changeFirst(slice []int) {
	slice[0] = 1000
}

func mutables() {
	//Mutalbes
	a := []int{1, 2, 3, 4}
	b := a
	b[0] = 100
	fmt.Println(a, b)

	// var c map[int]string = map[int]string{1: "Anirudh"}
	c := map[int]string{1: "Anirudh"}
	d := c
	d[2] = "Google"
	d[1] = "Anirudh Khurana"
	fmt.Println(c, d)

	// Imutables
	e := [5]int{1, 2, 3, 4, 5}
	f := e
	f[0] = 55
	fmt.Println(e, f)

	// Random Example
	g := [5]int{11, 22, 33, 44, 55}
	h := g[1:3]
	i := g
	j := h
	i[0] = 111
	j[0] = 1111
	fmt.Println(g, h, i, j)

	//Using Function
	k := []int{66, 77, 88}
	fmt.Println(k)
	changeFirst(k)
	fmt.Println(k)
}
