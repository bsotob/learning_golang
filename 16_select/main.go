//ultima parte del curso de uso de concurrencia
//El select permite a las gorutinas establecer multiples conexiones
//for + select (select permite parar un flujo, como un switch, nos permite parar el thread y coger valores de entradas/salida)
package main

import "fmt"

type sumaElements struct {
	a int
	b int
}

func sumaProceso(sumaElementos chan sumaElements, resultados chan int, quit chan int) {
	//El select bloquea el thread (como un swich,)
	for {
		select { //hasta que no reciba un sumaElementos se bloquea aquí.
		case sumaArealizar := <-sumaElementos:
			resultados <- sumaArealizar.a + sumaArealizar.b
		case <-quit:
			fmt.Println("Gorutina parada")
		}

	}
}

func main() {
	/*
		|
		|	\
		|->	| //enviamos sumaElementos
		|	|
		|<-	| //Nos retorna el valor
		|	| sumaProceso
		|
		Main
	*/
	canalElementos := make(chan sumaElements)
	canalResultados := make(chan int, 2) //50 de buffer porque en el bucle utilizamos 50 ¿intentos?
	canalQuit := make(chan int)
	go sumaProceso(canalElementos, canalResultados, canalQuit)
	canalElementos <- sumaElements{1, 2}
	i := 0

FR: //esto es un tag, con esto informamos que al hacer el break lo haremos de todo el for
	for {
		select {
		case resultado := <-canalResultados:
			i++
			fmt.Println(resultado)
			if i == 100000 {
				canalQuit <- 50
				//break //paramos ejecucion del select, no del for,
				break FR
			} else {
				canalElementos <- sumaElements{i, 2}
			}
		}
	}

}
