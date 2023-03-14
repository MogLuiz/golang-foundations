package main

// Um ponteiro é um endereçamento na memória

// Uma variável é um local que aponta para um ponteiro que tem um valor na memória

func main() {
	// Mémoria -> Endereço -> Valor

	// Variável -> Ponteiro que tem um endereço na memoria -> Valor
	a := 10

	// Dessa forma estou armazenando o endereçamento de memória a variável "a" e não o valor armazenado na mesma.
	var pointer *int = &a
	println(pointer)

	// Estou alterando o valor desse endereço de memoria. Dessa forma todas variaveis que apontar para o mesmo vão refletir essa alteração.
	*pointer = 20

	println(*pointer)
}
