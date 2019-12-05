package main

import "fmt"

func main() {
	//un map es un diccionario "mapa"
	//var m map[string]string
	var asignaturas map[string]string
	fmt.Println(asignaturas)
	deportes := map[string]int{
		"tiro con arco": 20,
		"ajedrez":       14,
		"jabalina":      5,
	}
	fmt.Println(deportes)

	fmt.Printf("Participantes de jabalina:  %v \n", deportes["jabalina"])
	resultado, ok := deportes["jabalina"]
	fmt.Println(ok)
	fmt.Println(resultado)
	if ok {
		fmt.Println("extraccion correcta")
	} else {
		fmt.Println("extraccion incorrecta")
	}
	//para eliminar un dato en un map
	fmt.Println(deportes)
	//delete(deportes, "tiro con arco")
	fmt.Println(deportes)
	//recorrer diccionarios
	for llave, valor := range deportes {
		if llave == "ajedrez" {
			fmt.Println(llave, valor)
		}
	}

}
