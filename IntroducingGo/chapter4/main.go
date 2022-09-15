package main

import "fmt"

func main() {
	// for loop
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("I love you evenly")
		} else {
			fmt.Println("I love you oddly")
		}
	}

	// switch stmt
	flag := 6
	switch flag {
	case 1:
		fmt.Print(flag)
	case 5:
		fmt.Println(flag)
	default:
		fmt.Println("none of these")
	}
}
