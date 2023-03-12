package main

import "fmt"

func main() {
	fmt.Println(sumWithVariadicFunction(100, 400, 600, 7000))
}

func sumWithVariadicFunction(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
