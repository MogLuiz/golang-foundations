package main

// Dentro do GO só temos uma forma de trabalhar com looping que é UTILIZANDO O "FOR"

func main() {

	// Basic FOR
	// for i := 0; i < 13; i++ {
	// 	println(i)
	// }

	numbers := []string{"one", "two", "tree"}

	// Basicamente o RANGE está referenciando meu slice
	for index, currentValue := range numbers {
		println(index, currentValue)
	}

	// Iterando escondendo o INDICE
	for _, currentValue := range numbers {
		println(currentValue)
	}

	// Utilizando o FOR como se fosse o WHILE
	i := 0
	for i < 10 {
		println(i)
		i++
	}

	// Loop infinito
	for {
		println("Loop infinito")
	}

}
