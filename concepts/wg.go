package concepts

import (
	"fmt"
	"sync"
)

func Wg_example() {
	// WaitGroup is a synchronization primitive that lets you
	fmt.Println("Hello, 世界")
	// done := make(chan bool)
	values := []string{"a", "b", "c"}

	var wg sync.WaitGroup
	for _, v := range values {
		fmt.Println(v, &v) // same instance of the v variable is used, hence the same address is printed

		wg.Add(1)
		v := v
		go func() {
			defer wg.Done()
			fmt.Println(v)
			// done <- true
		}()
		// }(v)
	}

	// wait for all goroutines to complete before exiting
	// for _ = range values {
	// 	<-done
	// }
	wg.Wait()
}

/* output
a 0xc00009e210
b 0xc00009e210
c 0xc00009e210
*/
/*
c a b or a b c - any order can come since the goroutines are async
*/
