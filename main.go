package main

import (
	"fmt"
)

func main() {
	/*	Un slice tiene tamaño dinamico son punteros a estructuras dinamucas.
		 Es una estructura con 3 valores:
			1- Un arrego comun de tamaño fijo
			2- Capacity
			3- Lenght
	*/

	// var l []int
	l := make([]int, 100)
	// El := define a 'l' con el tipo de lo que devuelve el miembro de la izquierda
	// El make devuelve un slice con la capacidad indicada, al superar la capacidad, vuelve a generar un sline nuevo
	l = append(l, 10)
	// El 'apprend esta creando un nuevo slice con todos los elementos de l agregando el valor
	fmt.Printf("%p\n", l) // Aca mostramos la direccion de memoria

	l = append(l, 100)
	fmt.Printf("%p\n", l)

	l = append(l, 1000)
	fmt.Printf("%p\n", l)

	fmt.Println(l)

	for index, value := range l {
		fmt.Println(index, value)
	}
}
