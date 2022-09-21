package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Contains("test", "es"))
	fmt.Println(strings.Count("test", "t"))
	fmt.Println(strings.HasPrefix("test", "te"))
	fmt.Println(strings.HasSuffix("test", "st"))
	fmt.Println(strings.Index("test", "e"))

	// join
	fmt.Println(strings.Join([]string{"hello", "world"}, "-"))

	// repeat
	fmt.Println(strings.Repeat("iloveyou", 10))

	// split - returns an array of splitted strings
	fmt.Println(strings.Split("a-b-c-d-e", "-"))

	// tolower
	fmt.Println(strings.ToLower("Hello WOrld"))

	// toupper
	fmt.Println(strings.ToUpper("hello world"))

	arr := []byte("test")
	fmt.Println(arr)
}
