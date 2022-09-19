package main

import "fmt"

// normal func
func avg(arr []float64) float64 {
	total := 0.0
	for _, val := range arr {
		total += val
	}
	return total / float64(len(arr))
}

// func - return types
func greet() (ans string) {
	ans = "Good morning!!!"
	return
}

// func - return multiple values
func scores() (int, int) {
	return 9, 10
}

// variadic functions
func add(args ...int) int {
	total := 0
	for _, val := range args {
		total += val
	}
	return total
}

// recursion
func fact(x uint) uint {
	// base
	if x == 0 {
		return 1
	}
	// main
	return x * fact(x-1)
}

// panic
func dontShoot() {
	defer func() {
		str := recover()
		fmt.Println(str, "-- called inside recover")
	}()
	panic("Please don't shoot!!!")
}

func first() {
	fmt.Println("File opened!!!")
}

func second() {
	fmt.Println("File closed at last!!!")
}

func openClose() {
	defer second()
	first()
	panic("ff") // defer called before panic even
}

func main() {
	arr := []float64{1, 2, 3, 4, 5}
	fmt.Println("Avg of the arr is", avg(arr))
	fmt.Println("Greet: ", greet())

	x, y := scores() // assign returned values from the function
	fmt.Println("Scores: ", x, y)

	// variadic func
	fmt.Println(add(1, 2, 3, 4, 5))
	xs := []int{5, 4, 3, 2, 1}
	fmt.Println(add(xs...)) // slice unpacking

	// factorial
	fmt.Println(fact(5))

	// defer
	// openClose()

	// panic
	dontShoot()
}
