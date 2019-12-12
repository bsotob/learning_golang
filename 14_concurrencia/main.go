package main

import (
	"fmt"
	"time"
)

//gorutine "simula lo que haría un thread."
//clouseres y callbacks los utilizamos para que un proceso reciba una función al acabar

func run(num int, callback func(int)) {
	for i := 0; i < 10; i++ {
		fmt.Printf("Gorutina : %v %v\n", num, i)
	}
	callback(num)
}

func main() {
	timeStart := time.Now()
	numeroThreads := 50
	callback := func(numGoturina int) {
		numeroThreads--
		fmt.Printf("Ha finalizado la rutina: %v\n", numGoturina)
	}
	for i := 0; i < 10; i++ {
		go run(i, callback)
	}
	//for i := 0; i < 10; i++ {
	//	fmt.Println(i)
	//}
	for {
		if numeroThreads < 100 {
			fmt.Println("no hay threads activos")
			break
		}
	}
	timeFinal := time.Now()
	fmt.Println(timeFinal.Sub(timeStart))
}
