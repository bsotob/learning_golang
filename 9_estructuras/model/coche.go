package model

type Motor struct {
	NumCilindros int
	Cilidrada    int
}
type Coche struct {
	MotorCoche   *Motor
	NumeroRuedas int
	//radio	*bool
	LucesAntiniebla *int //posibilidad de tener un entero, que puede permitirse ser nulo (puede tener punteros)
	numeSerie       string
}

//creamos constructor
func NewCoche(numCilindros int, cilindrada int, numRuedas int) Coche {
	var motor *Motor
	if numCilindros > 0 {
		motor = &Motor{
			NumCilindros: numCilindros,
			Cilidrada:    cilindrada,
		}
	}
	return Coche{
		MotorCoche:   motor,
		NumeroRuedas: numRuedas,
		numeSerie:    "asdfg",
	}
}

//para declarar una struc :
//type <nombre de la struct> struct {}
//para poner las variables es necesario tanto declarar con nombre como el tipo
// <nombre> <tipo>
