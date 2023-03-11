package main

import "fmt"

type ID int

var (
	a ID = 2
)

func main() {
	fmt.Printf("O tipo de 'A' é %T", a)
	fmt.Printf("O valor de 'A' é %v", a)
}
