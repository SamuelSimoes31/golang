package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	wg.Add(5)
	for i := 0; i < 5; i++ {
		// x := i
		go func() {
			wg.Done()
			fmt.Println(i)
		}()
	}

	wg.Wait()
}
