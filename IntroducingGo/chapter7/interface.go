package main

import (
	"fmt"
	"math"
)

// interface
type Shape interface {
	area() float64
}

// Circle
type Circle struct {
	r float64
}

func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

// Rectangle
type Rectangle struct {
	l, b float64
}

func (r *Rectangle) area() float64 {
	return r.l * r.b
}

// func accepting an interface
func calcArea(shape Shape) float64 {
	return shape.area()
}

func main() {
	c1 := &Circle{10}
	r1 := &Rectangle{10, 10}
	fmt.Println("Circle area is ", calcArea(c1))
	fmt.Println("Rectangle area is", calcArea(r1))
}
