package main

import "fmt"

type Point struct {
	xx int
	yy int
}

type Circle struct {
	radius int
	center Point
	//*Point			We can even write like this
}

func structs2() {
	x := Circle{25, Point{5, 6}}
	fmt.Println(x, x.center.xx)

}
