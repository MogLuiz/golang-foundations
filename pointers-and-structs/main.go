package main

// Trabalhamos com ponteiros quando necessitamos gerar mutações nos valores recebidos por parâmetro em todo o projeto.

type Account struct {
	balance int
}

func (a *Account) simulate(value int) int {
	a.balance += value
	println(a.balance)
	return a.balance
}

func NewAccount() *Account {
	return &Account{balance: 0}
}

func main() {
	bank := Account{
		balance: 100,
	}
	bank.simulate(200)
	println(bank.balance)
}
