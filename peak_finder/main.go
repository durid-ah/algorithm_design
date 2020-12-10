package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello")
	s := 1
	b := &s 
	fmt.Println(*b)
}