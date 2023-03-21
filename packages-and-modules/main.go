package main

import (
	"fmt"

	"localPackages/math"
)

func main() {
	sumResult := math.Sum(10, 20)
	car := math.Carro{Marca: "bmw"}

	fmt.Println(car.Marca)
	fmt.Println("Resultado da soma: ", sumResult)
}
