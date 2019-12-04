package model

import "fmt"

//modelo privado
type motor struct {
	NumCilindros int
	Cilidrada    int
}

//modelo publico
type Coche struct { //con minuscula, hacemos que la clase sea privada
	MotorCoche   *motor
	numeroRuedas int
	//radio	*bool
	LucesAntiniebla *int //posibilidad de tener un entero, que puede permitirse ser nulo (puede tener punteros)
	numeSerie       string
}

//creamos constructor
func NewCoche(numCilindros int, cilindrada int, numRuedas int) Coche {
	var motorCoche *motor
	if numCilindros > 0 {
		motorCoche = &motor{
			NumCilindros: numCilindros,
			Cilidrada:    cilindrada,
		}
	}
	return Coche{
		MotorCoche:   motorCoche,
		numeroRuedas: numRuedas,
		numeSerie:    "asdfg",
	}
}
func (c Coche) Arranca() string { //creamos un método para la estructura

	return fmt.Sprintf("El motor ha arrancado correctamente con %v de cilidrada", c.MotorCoche.Cilidrada)
}

//para declarar una struc :
//type <nombre de la struct> struct {}
//para poner las variables es necesario tanto declarar con nombre como el tipo
// <nombre> <tipo>
