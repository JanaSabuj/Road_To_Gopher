package concepts

import (
	"fmt"
	"sync"
)

type Pod struct {
	Name string
	UUID string
}

// Pattern: Deleting an array of k8s pods concurrently
func Wg_example() {
	fmt.Println("Hello, 世界")

	// slice of Pods
	pods := []Pod{
		{Name: "Pod 1", UUID: "uuid1"},
		{Name: "Pod 2", UUID: "uuid2"},
		{Name: "Pod 3", UUID: "uuid3"},
	}
	fmt.Println(pods)

	var wg sync.WaitGroup
	for _, p := range pods {
		// 	fmt.Println(v, &v) // same instance of the v variable is used, hence the same address is printed
		wg.Add(1)
		p := p
		go func() {
			defer wg.Done()
			fmt.Printf("Pod %v is deleted\n", p.Name)
			// done <- true
		}()
	}

	wg.Wait() // wait till all Pods are deleted
	fmt.Println("All pods have been deleted")
}

/* output
a 0xc00009e210
b 0xc00009e210
c 0xc00009e210
*/
/*
c a b or a b c - any order can come since the goroutines are async
*/
