package main

import (
	"fmt"
	"like/grammar"
)

func main() {
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
