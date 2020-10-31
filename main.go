package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[0] = "a"
	m[1] = "b"

	for key, val := range m {
		fmt.Println(key, val)
	}
}
