package main

import (
	"fmt"
	"math"
)

// struct
type Circle struct {
	x, y, r float64
}

func circleArea(c *Circle) float64 {
	return math.Pi * c.r * c.r
}

// method of a struct
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c1 := Circle{x: 10, y: 20, r: 30}
	fmt.Println(c1) //

	c2 := &Circle{1, 2, 3}
	fmt.Println(c2)  // pointer
	fmt.Println(*c2) // dereference a pointer

	fmt.Println(c2.x, c2.y, c2.r)
	fmt.Println(circleArea(&Circle{0, 0, 3})) // passing a struct to a function
	fmt.Println(c2.area())                    // calling the method of a struct
}
