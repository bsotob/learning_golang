package main

import "fmt"

func main() {
	//Arrays con numero de variables definido
	//contexto uso: si sabemos cuantos elementos tendrá el array
	var hola [2]string
	hola[0] = "hola"
	hola[1] = "mundo"
	fmt.Println(hola)
	//cuando no sabemos cuantos numeros tendrán
	var numeros []int
	fmt.Println(numeros)
	for i := 0; i < 20; i++ {
		numeros = append(numeros, i) //con la funcion appen añadimos datos al array, realizamos copia de la var array.
	}
	fmt.Println(numeros)
	//si queremos evitar hacer "var numeros [] int y utilizar :=" ->
	//letras := []string{"uno","dos"}
	letras := []string{}
	letras = append(letras, "hola", "adios", "ye")
	fmt.Println(letras)
	//recorremos array por elemento
	for id, campos := range letras { //si queremos borrar el id (indice del array, pondriamos "_,campos")
		fmt.Println(id, campos)
	}
	for _, campos := range letras { //si queremos borrar el id (indice del array, pondriamos "_,campos")
		fmt.Println(campos)
	}
	//funcion longitud de un array es -> len()
	longitudArray := letras
	fmt.Println(longitudArray)
	//para saber la capacidad de un array -> cap()
	capacidadArray := cap(letras)
	fmt.Println(capacidadArray)
}
