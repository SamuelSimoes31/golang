package main

import (
	"fmt"
	"runtime"
)

func main() {
	cpus_logicas := runtime.NumCPU()
	fmt.Println(cpus_logicas)
}

/*
Cores lógicos = x *y
• x é número de cores físicos
• y é o número de threads de
hardware por core físico
*/
