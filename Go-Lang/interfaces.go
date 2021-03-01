package main

import "fmt"

type circle struct {
	radius float64
}

type rectangle struct {
	length  float64
	breadth float64
}

func (c circle) area() float64 {
	return c.radius * c.radius * 3.14
}

func (r rectangle) area() float64 {
	return r.length * r.breadth
}

type shape interface {
	area() float64
}

func interfaces() {
	c1 := circle{3.4}
	r1 := rectangle{3.4, 6.7}

	sha := []shape{&c1, &r1}
	for _, v := range sha {
		fmt.Println(v.area())
	}
}
