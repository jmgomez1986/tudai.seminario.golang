package main

import "fmt"

type person struct {
	name string
}

type figure struct {
	name string
}

// Printable ...
type Printable interface {
	print()
}

func (p *person) print() {
	fmt.Println("[person]", p.name)
}

func (f *figure) print() {
	fmt.Println("[figure]", f.name)
}

func invokePint(p Printable) {
	p.print()
}

func main() {
	p := &person{name: "Juan"}
	f := &figure{"Circulo"}

	invokePint(p)
	invokePint(f)

}
