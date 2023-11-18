package main

import (
	"fmt"
	"sync"
)

func f1(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Println("f1:", i)
	}
}

func f2(wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 20; i++ {
		fmt.Println("   f2:", i)
	}
}

func main() {
	wg := sync.WaitGroup{}

	wg.Add(1)
	go f1(&wg)
	wg.Add(1)
	go f2(&wg)

	wg.Wait()

}