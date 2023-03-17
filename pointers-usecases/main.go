package main

// Quando trabalhamos com GO, toda vez que passamos um parâmetro, estamos mandando uma cópia desse valor
// que está na memória. Não estamos passando exatamente o valor qu está na memória.

// Nessa função estou recebendo a cópia dos valores. E não o endereço da memoria do valor em sí
func soma1(a, b int) int {
	// Dessa forma, atribuindo valor a variável "a" estou alterando apenas a cópia
	a = 50
	return a + b
}

// Para alterarmos os parâmetros recebidos de forma global e não apenas a "cópia" do valor precisamos passar o
// endereço de memória daquela variável.
// PARA ISSO UTILIZAMOS OS PONTEIROS.

// Nessa função estou recebendo a referência do endereço de memória dessas variáveis.
func soma2(a, b *int) int {
	// Dessa forma, atribuindo valor a variável "a" estou alterando de forma global.
	// Pois estou alterando a referência de memória dessa variável. E isso reflete de forma global.
	*a = 50
	return *a + *b
}

func main() {
	var1 := 10
	var2 := 20

	// Passando uma cópia dos valores.
	soma1(var1, var2)
	println(var1, var2)

	// Passando a referência da memória. usamos o "&"
	soma2(&var1, &var2)
	println(var1, var2)
}
