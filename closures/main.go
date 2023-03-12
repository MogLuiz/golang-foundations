package main

import "fmt"

func main() {
	total := func() int {
		return sumWithVariadicFunction(12, 20) * 2
	}()

	fmt.Println(total)
}

func sumWithVariadicFunction(numbers ...int) int {
	total := 0
	for _, number := range numbers {
		total += number
	}
	return total
}
