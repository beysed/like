package main

import (
	"fmt"
	"like/grammar"
)

type A struct {
	my string
}

func main() {
	a := A{my: "a"}
	fmt.Println(a.my)

	res, err := grammar.Parse("exmample.like", []byte("hello"))

	fmt.Printf("Hello, World\n")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Printf("Result: %v\n", res)
	}

	if res != nil {
		fmt.Printf("Result: %v\n", res)
	}
}
