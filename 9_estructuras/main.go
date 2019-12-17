package main

import (
	"fmt"
	"learning_golang/9_estructuras/model"
)

func main() {
	coche := model.NewCoche(10, 1000, 1)
	fmt.Println(coche.Arranca())
	fmt.Println(coche.MotorCoche)

	//
}
