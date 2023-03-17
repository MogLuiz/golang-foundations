package main

import "fmt"

type Customer struct {
	name string
}

func (c Customer) walk() {
	c.name = "Luiz Henrique"
	fmt.Printf("O cliente %v andou\n", c.name)
}

func main() {
	Luiz := Customer{
		name: "Luizin",
	}
	Luiz.walk()
	fmt.Printf("O cliente %v\n", Luiz.name)
}
