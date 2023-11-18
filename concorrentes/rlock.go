package main

import (
	"fmt"
	"sync"
)

func main() {
	var mu sync.RWMutex
	var balance int

	mu.RLock()           //1
	fmt.Println(balance) //1
	mu.RLock()           // 2
	fmt.Println(balance) // 2
	mu.RUnlock()         // 2
	mu.RUnlock()         //1
	mu.Lock()            //  3
	balance++            //  3
	mu.Unlock()          //  3
	mu.Lock()            //   4
	balance++            //   4
	mu.Unlock()          //   4

	mu.RLock()
	fmt.Println(balance)
	mu.RUnlock()
}
