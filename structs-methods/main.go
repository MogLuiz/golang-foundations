package main

import "fmt"

type Address struct {
	street string
	number int
	city   string
	state  string
}

type Customer struct {
	name   string
	age    int
	active bool
	Address
}

func (c Customer) isActiveCustomer() string {
	if c.active == true {
		return "ativo"
	}
	return "inativo"
}

func main() {
	Luiz := Customer{
		name:   "Luiz",
		age:    20,
		active: true,
	}

	fmt.Printf("Cliente: %s, possui %d anos de idade e o seu status no sistema é: %s\n", Luiz.name, Luiz.age, Luiz.isActiveCustomer())
}
