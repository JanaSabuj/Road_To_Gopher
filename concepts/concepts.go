package concepts // pkg names are in small

import (
	"fmt"
)

func foo() error {
	// return errors.New("foo error")
	return nil
}

// func having multiple return values
func foo1() (string, error) {
	// return errors.New("foo error")
	return "empty", nil
}

func loop() {
	// for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	// while loop
	sum = 0
	for sum < 10 {
		sum += 1
	}
	fmt.Println(sum)

	// infinite loop
	sum = 0
	for {
		sum += 1
		if (sum > 20){
			break
		}
	}
	fmt.Println(sum)
}

func deferEx() {
	defer func() {
		fmt.Println("defer func called... exiting")
	}()

	fmt.Println("func working....")
}

func newMakeInit() {
	b := Book{}
	fmt.Println("init book", b)

	var b1 Book
	fmt.Println("var book", b1)

	b2 := new(Book) // return a ptr to zero value struct of the same
	fmt.Println("new book", b2)

	slice := make([]int, 10)
	fmt.Println(slice)
}

type Sabuj struct {
	Name string
}

// actual mutation by ref
func (s *Sabuj) setNameRef(name string) {
	s.Name = name
}

// NO actual mutation
func (s Sabuj) setNameVal(name string) {
	s.Name = name
}

func (s *Sabuj) printName() {
	fmt.Println(s)
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

	// for, while, infinite loops
	loop()

	// defer
	deferEx()

	// new vs make vs initialisation
	newMakeInit()

	// _ -> ignore a reutrn value
	s := Sabuj{}
	s.setNameRef("Harry")
	s.printName()
	s.setNameVal("Potter") // no actual mutation takes place
	s.printName()

	// scoping of go vars
	// Caps is exported out of package
	// small letters is not exported outta package

}
