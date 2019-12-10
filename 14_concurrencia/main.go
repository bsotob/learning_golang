package main

import "fmt"

//gorutine "simula lo que haría un thread."

func run() { //para lanzar threads se suele utilizar la palabra RUN para llamar a la funcion
	for i := 0; i < 10; i++ {
		fmt.Printf("Gorutina %v\n", i)
	}

}

func main() {
	go run() //con la funcion go lanzamos una gorutina
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
	for {

	}
}
