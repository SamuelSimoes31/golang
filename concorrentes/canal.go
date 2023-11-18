package main

import "fmt"

func gorotina01(ch chan int){
	ch <- 1
}

func gorotina02(ch chan int){
	i := <- ch
	fmt.Println("Inteiro recebido: ", i)
}


func main() {
	ch := make(chan int)

	go gorotina01(ch)
	go gorotina02(ch)

	fmt.Scanln()
}