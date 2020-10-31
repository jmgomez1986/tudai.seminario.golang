package main

import "fmt"

type close func(s string)

func invokeClose(c close) {
	c("Hello world")
}

func main() {
	f := func(s string) {
		fmt.Println(s)
	}

	invokeClose(f)

}
