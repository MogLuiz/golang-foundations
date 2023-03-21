package main

func SumInteger(m map[string]int) int {
	var sum int
	for _, v := range m {
		sum += v
	}

	return sum
}

type MyNumber int

type Number interface {
	~int | float64
}

// Dessa forma eu criei uma função genérica que aceita tanto Inteiros quanto Float64
func SumWithGeneric[T Number](m map[string]T) T {
	var sum T
	for _, v := range m {
		sum += v
	}

	return sum
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	mapInteger := map[string]int{"Luiz": 1000, "Lucas": 2000}
	mapGeneric := map[string]float64{"Luiz": 1000.9, "Lucas": 2000.3}
	map3 := map[string]MyNumber{"Luiz": 1000, "Lucas": 2000}
	println(SumInteger(mapInteger))
	println(SumWithGeneric(mapGeneric))
	println(SumWithGeneric(map3))
}
