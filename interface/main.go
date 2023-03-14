package main

import "fmt"

type Address struct {
	street string
	number int
	city   string
	state  string
}

// Customer é um Person porque ele tá implementando o método Deactivate() na linha 24 que é da interface Pessoa.
type Person interface {
	Desactivate()
}

type Customer struct {
	name   string
	age    int
	active bool
	Address
}

func (c Customer) Desactivate() {
	c.active = false
	fmt.Printf("O cliente %s foi desativado\n", c.name)
}

func Desactivation(person Person) {
	person.Desactivate()
}

func main() {
	Luiz := Customer{
		name:   "Luiz",
		age:    20,
		active: true,
	}

	Luiz.Desactivate()

	fmt.Printf("Cliente: %s, possui %d anos de idade.", Luiz.name, Luiz.age)
}
