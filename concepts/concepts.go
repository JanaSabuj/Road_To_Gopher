package concepts // pkg names are in small

import (
	"fmt"
)

func foo() error {
	// return errors.New("foo error")
	return nil
}

func foo1() (string, error) {
	// return errors.New("foo error")
	return "empty", nil
}

func Concepts() {  
	fmt.Println("Welcome to concepts.go")

	// mixedCase/camelcase naming convention (snake_case, kebab-case, PascalCase, )
	// var mixedCase string 
	mixedCase := "camelCase"
	fmt.Println(mixedCase)

	var err error
	// if and switch statements can take initialisation blocks
	if err = foo(); err != nil { 
		fmt.Println("Error is not nil", err)
	}

	a, err := foo1() // you can use := syntax if atleast 1 new varibale is on the LHS, the rest would be reassigned
	fmt.Println(a, err)
	// err := foo() X
	b, err := foo1()
	fmt.Println(b, err)

}
