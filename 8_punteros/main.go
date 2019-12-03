package main

import "fmt"

func main() {
	//en go no existe el nulo en variables
	//el puntero es la dirección de memoria de una variable
	a := 2
	pA := &a //esto es una dirección de memoria
	fmt.Println("dirección de memoria de a ->)", pA)
	fmt.Println(pA)
	fmt.Println(&a)
	fmt.Println(*pA) //accede dentro de pA (por lo que nos dara valor de a

	//si pasamos un puntero (a una funcion) podemos el contenido del puntero.
	//siempre hay que chequear si un puntero es nulo, si es nulo dará un panic
	//cuando trabajamos con terceras librerias hay q revisar los valores para q no devuelvan nil.
	pNill := returnNill()
	if pNill == nil {
		fmt.Println("no podemos ejecutar, porque saldrá panic")
	}
	fmt.Println(a)
	main2(pA)
	fmt.Println("funcion main2")
	fmt.Println(a)
	fmt.Println(returnNill())
}
func main2(pA *int) {
	*pA = 4
}
func returnNill() *int {
	return nil //nil es null
}
