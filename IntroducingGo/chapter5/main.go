package main

import (
	"fmt"
)

func main() {
	// arrays - fixed length
	var x [4]int
	x[3] = 100
	fmt.Println(x)

	// array1 - fixed length
	arr := [5]float64{1, 2, 3, 4, 5}
	fmt.Println("Arr is", arr)
	total := 0.0
	for _, val := range arr {
		total += val
	}

	fmt.Println("Total is", total)
	fmt.Println("Avg is", total/float64(len(arr)))

	// slice - dynamic array
	slice := []int{1, 2, 3}
	slice = append(slice, 4, 5)
	fmt.Println(slice)

	// copy - takes two args src & dest
	s1 := []int{1, 2, 5, 6, 7}
	s2 := make([]int, 3)
	copy(s2, s1)
	fmt.Println("After copying", s1, s2)
}
