package main

import (
	"fmt"

	"github.com/SKAhack/go-shortid"
)

func main() {
	g := shortid.Generator()
	for i := 0; i < 10; i++ {
		fmt.Println(g.Generate())
	}
}
