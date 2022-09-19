package main

import "fmt"

func zero(x *int) {
	*x = 0
}

func one(x *int) {
	*x = 1
}

func main() {
	x := 4
	zero(&x)
	fmt.Println(x)

	// new - allocates a pointer
	y := new(int)
	one(y)
	fmt.Println(*y)
}
