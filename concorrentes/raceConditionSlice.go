package main

import "fmt"

func main() {
	var x []int

	go func() {
		fmt.Println("x[9999] = 1")
		x[9999] = 1
	}()

	go func() {
		fmt.Println("make 10")
		x = make([]int, 10)
	}()

	go func() {
		fmt.Println("make 1000000")
		x = make([]int, 1000000)
	}()

	fmt.Scanln()

}
