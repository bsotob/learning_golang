package main

import (
	"curso_go/9_estructuras/model"
	"fmt"
)

func main() {
	coche := model.NewCoche(10, 1000, 4)
	fmt.Println(coche.NumeroRuedas)
	fmt.Println(coche.MotorCoche.Cilidrada)
}
