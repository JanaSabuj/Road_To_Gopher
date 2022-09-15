package main

import "fmt"

func main() {
	// strings
	var x string = "Hello World" // normal
	var y = "Hi"                 // infer type from var stmt
	z := "Hey there Delilah"     // compiler can infer the type based on the literal you assign
	z = "NYC"
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	dogsName := "Whiskers"
	fmt.Println("My dog's name is", dogsName)

}
