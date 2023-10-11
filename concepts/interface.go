package concepts

import "fmt"

// interface - defines a common behaviour of execution
type Animal interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d *Dog) Speak() string {
	return "Bark!"
}

func (d *Dog) Protect() string {
	return "Hooman!"
}

type Human struct{}

func (h Human) Speak() string {
	return "Hello!"
}

func Printall(vals []interface{}) {
	for _, val := range vals {
		fmt.Println(val)
	}
}

func Intf() {
	animals := []Animal{
		&Dog{Name: "TOMMY"}, // & given because interface function is implemented by a pointer reciever
		&Human{},            // pointer value can resolve to normal value also, but vice versa is not true
	}

	for _, animal := range animals {
		fmt.Println(animal.Speak())

		// casting an interface to its underlying type
		if v, ok := animal.(*Dog); ok {
			fmt.Println(v.Protect(), v.Speak())
		}
	}

	names := []string{"Harrry", "Potter", "Ron"}
	vals := make([]interface{}, len(names))
	for idx, n := range names {
		vals[idx] = n
	}

	Printall(vals)
}

/*
+------------+
| *itab      |  (Type Pointer)
+------------+
| *data      |  (Data Pointer)
+------------+
*/
