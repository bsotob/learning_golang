package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	file_data, err := ioutil.ReadFile("./ficheros/hola.txt")

	if err != nil {
		fmt.Println("hubo un error al cargar el archivo")
	}
	fmt.Println(string(file_data))
}
