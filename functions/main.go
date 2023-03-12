package main

import (
	"errors"
	"fmt"
)

func main() {
	value, err := sum(20, -30)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(value)

}

func sum(a, b int) (int, error) {
	result := a + b
	if result >= 0 {
		return result, nil
	}
	return result, errors.New("Valor da soma é negativo.")
}

func sum1(a int, b int) int {
	return a + b
}

// Quando os parâmetros de uma função possui os mesmos tipos, pode ser declarado uma unica vez no último parâmetro
func sum2(a, b int) int {
	return a + b
}
