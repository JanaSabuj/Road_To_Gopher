package concepts

import (
	"fmt"
	"unsafe"
)

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
}
