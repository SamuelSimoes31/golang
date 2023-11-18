package main

import (
	"fmt"
	"time"
)

func main() {
	chn1 := make(chan string)
	chn2 := make(chan string)

	go func(){
		for {
			chn1 <- "from 1"
			time.Sleep(time.Second * 2)
		}
	}()

	go func(){
		for {
			chn2 <- "from 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func(){
		for {
			select {
			case msg1 := <- chn1:
					fmt.Println(msg1)
			case msg2 := <- chn2:
					fmt.Println(msg2)
			}
		}
	}()

	fmt.Scanln()
}