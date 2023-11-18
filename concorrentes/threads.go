package main

import "fmt"

func f1() {
	for i := 0; i < 20; i++ {
		fmt.Println("f1:", i)
	}
}

func f2() {
	for i := 0; i < 20; i++ {
		fmt.Println("f2:", i)
	}
}

func main() {
	go f1()
	go f2()

	_, _ = fmt.Scanln()
}