package main

import "fmt"

type person struct {
	name string
}

func (p *person) print() {
	fmt.Println(p.name)
}

func (p *person) clean() {
	p.name = ""
}

// Poniendo el '*' usa un puntero a la misma ponsicion de memoria del que lo llamo, si no hace una copia

func main() {
	p := &person{name: "Juan"}
	// Crea una instancia de la estructura 'Person',
	//le agrega valor al atributo y obtiene el valor del puntero y lo guarda en 'p'

	p.print()
	p.clean()
	p.print()

}
