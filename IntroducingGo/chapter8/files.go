package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	bs, err := ioutil.ReadFile("IntroducingGo/chapter8/abc.txt") // read file as ascii values
	fmt.Println(bs, err)
	str := string(bs)
	fmt.Println(str)

	// os
	// file, err := os.Create("test.txt")
	// if err != nil {
	// 	return 
	// }
}