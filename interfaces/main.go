package main

import "fmt"

type User interface {
	Permisos() int
	Nombre() string
}
type Admin struct {
	nombre string
}

func (this Admin) Permisos() int {
	return 5
}
func (this Admin) Nombre() string {
	return this.nombre
}
func auth(user User) string {
	if user.Permisos() > 4 {

		return "tiene permisos de administrador"
	}
	return ""
}
func main() {
	admin := Admin{"Borja"}
	fmt.Println(auth(admin))
	array_playbook := make([]User, 2)
	array_pla := append(array_playbook, admin)
	fmt.Println(array_pla)
}
