package main

import (
	"fmt"
)

func main() {
	var arr [2]int
	arr[0] = 1
	arr[1] = 2

	for index, value := range arr { // El 'range' devuelve cada una de las posiciones del arreglo
		fmt.Println(index, value)
	}
}
