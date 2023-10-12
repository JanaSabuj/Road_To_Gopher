package concepts

import (
	"fmt"
	"unsafe"
)

// Size of concepts.Employee struct: 56 bytes
type Employee struct {
	ID     int    // 8
	Name   string // 16
	Age    int16  // 2
	Gender string // 16
	Active bool   // 1
}

// type Employee struct {
// 	ID     int    // 8
// 	Name   string // 16
// 	Gender string // 16
// 	Age    int16  // 2
// 	Active bool   // 1
// }
// Size of concepts.Employee struct: 48 bytes

func Sz() {
	a := int(123)
	b := int64(123)
	c := "foo"
	d := struct {
		FieldA float32
		FieldB string
	}{0, "bar"}

	fmt.Printf("a: %T, %d\n", a, unsafe.Sizeof(a))
	fmt.Printf("b: %T, %d\n", b, unsafe.Sizeof(b))
	fmt.Printf("c: %T, %d\n", c, unsafe.Sizeof(c)) // string 16 bytes
	fmt.Printf("d: %T, %d\n", d, unsafe.Sizeof(d))

	var e Employee
	fmt.Printf("Size of %T struct: %d bytes", e, unsafe.Sizeof(e))
}

// https://perennialsky.medium.com/memory-allocation-for-a-struct-in-golang-b5057b8ccf23
