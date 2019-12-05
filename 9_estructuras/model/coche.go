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
	NumeroRuedas int
	//radio	*bool
	roto            bool
	LucesAntiniebla *int //posibilidad de tener un entero, que puede permitirse ser nulo (puede tener punteros)
	numeSerie       string
}

//creamos constructor
//en las estructuras siempre es bueno trabajar con punteros
func NewCoche(numCilindros int, cilindrada int, numRuedas int) *Coche {
	var motorCoche *motor
	if numCilindros > 0 {
		motorCoche = &motor{
			NumCilindros: numCilindros,
			Cilidrada:    cilindrada,
		}
	}
	return &Coche{
		MotorCoche:   motorCoche,
		NumeroRuedas: numRuedas,
		numeSerie:    "asdfg",
		roto:         false,
	}
}
func (c *Coche) Arranca() string { //creamos un método para la estructura
	if c.roto == false {
		return fmt.Sprintf("El motor ha arrancado correctamente con %v de cilidrada", c.MotorCoche.Cilidrada)
	} else {
		return fmt.Sprintf("No arrancamos, esta roto")
	}
}
func (c *Coche) IncrementarPotencia(addCilindrada int) string {
	c.MotorCoche.Cilidrada = c.MotorCoche.Cilidrada + addCilindrada
	return fmt.Sprintf("La nueva cilidrada es de %v", c.MotorCoche.Cilidrada)
}
func (c *Coche) PinchaRueda() string {
	if c.NumeroRuedas == 1 {
		c.NumeroRuedas = 0
		c.roto = false
		//return fmt.Println("ruedas pinchadas, ya no hay ruedas disponibles")
	} else {
		c.NumeroRuedas = c.NumeroRuedas - 1
		//return fmt.Sprintf("hemos pinchado rueda, quedan %v ruedas diponibles", c.NumeroRuedas)

	}
	return fmt.Sprintf("numero de ruedas restantes : %v", c.NumeroRuedas)
}

//para declarar una struc :
//type <nombre de la struct> struct {}
//para poner las variables es necesario tanto declarar con nombre como el tipo
// <nombre> <tipo>
