package main

import "fmt"

type Customer struct {
	name   string
	age    int
	active bool
}

func main() {
	Luiz := Customer{
		name:   "Luiz",
		age:    20,
		active: true,
	}
	Luiz.active = false

	fmt.Printf("Cliente: %s, possui %d anos de idade e o seu status no sistema é: %s\n", Luiz.name, Luiz.age, isActiveCustomer(Luiz.active))
}

func isActiveCustomer(active bool) string {
	if active == true {
		return "ativo"
	}
	return "inativo"
}
