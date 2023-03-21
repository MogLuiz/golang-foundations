package math

// Tudo que eu colocar com a primeira Letra maiuscula (Fun, Struct, Variable) está sendo exportado do meu package
// O que eu colocar com a primeira letra minuscula não está sendo exportado

// Está sendo exportado pois começa com letra maiuscula
func Sum[T int | float64](a, b T) T {
	return a + b
}

// Não está sendo exportado pois começa com letra minuscula
func test[T int | float64](a, b T) T {
	return a + b
}

type Carro struct {
	Marca string
}
