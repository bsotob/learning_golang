package main

import (
	"fmt"
	"learning_golang/9_estructuras/model"
)

func main() {
	coche := model.NewCoche(10, 1000, 4)
	fmt.Println(coche.NumeroRuedas)
	fmt.Println(coche.MotorCoche.Cilidrada)
}
