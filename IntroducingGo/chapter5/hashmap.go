package main

import "fmt"

func main() {
	// map - k-v pair
	mp := make(map[string]int)
	mp["foo"] = 10
	fmt.Println("Len of mp is", len(mp))

	mp["bar"] = 20
	fmt.Println("Len of mp is", len(mp))

	fmt.Println(mp)
	fmt.Println(mp["foo"])

	// key not present in map - returns zero value of value type
	fmt.Println(mp["xyz"])
	if value, ok := mp["foo"]; ok {
		fmt.Println(value, ok)
	}

	// shorter way to create maps
	elements := map[string]string{
		"H":  "Hydrogen",
		"Li": "Lithium",
		"Be": "Beryllium",
	}

	fmt.Println(elements)
}
