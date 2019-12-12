package main

import "fmt"

//un canal nos permite hablar con threads diferentes
//con esto evitamos callbacks
//funcion make, nos permite generar tipos complejos de información, instancia o define, lo utilizamos para definir

func suma(a, b int, canal chan int) {
	canal <- a + b //<- con esto introducimos dentro del canal el resultado de a + b

}

func main() {
	canal := make(chan int)
	go suma(10, 10, canal)
	go suma(10, 100, canal)
	resultado1, resultado2 := <-canal, <-canal //pausamos el thread principal hasta que yo reciba la información del canal.
	fmt.Println(resultado1, resultado2)

	//otra forma de crear un canal es utilizar un buffer

	//creamos canal buffer
	canalBuffer := make(chan int, 2)
	canalBuffer <- 1
	//TH ---------1-------------- TH
	canalBuffer <- 2
	//TH --------2----------------- Error si no empleamos un canal con buffer
	fmt.Println(<-canalBuffer, <-canalBuffer)
	close(canalBuffer)
	//ATENCION, hay que cerrar el canal siempre funcion close
	//creamos funcion
	for resultado := range canalBuffer {
		fmt.Println(resultado)
	}

}
