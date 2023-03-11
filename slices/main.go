package main

import "fmt"

func main() {
	slice := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}

	fmt.Printf("length=%d capacity=%d %v\n", len(slice), cap(slice), slice)

	// Excluindo tudo do slice a partir do index 0
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[:0]), cap(slice[:0]), slice[:0])

	// Excluindo tudo do slice a partir da index 4
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[:4]), cap(slice[:4]), slice[:4])

	// Excluindo items do slice anteriores ao index 2
	fmt.Printf("length=%d capacity=%d %v\n", len(slice[2:]), cap(slice[2:]), slice[2:])

	// Adicionando o 110 na última posição do meu slice
	slice = append(slice, 110)
	fmt.Printf("length=%d capacity=%d %v\n", len(slice), cap(slice), slice)
}
