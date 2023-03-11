package main

import "fmt"

func main() {
	var firstArray [3]int
	firstArray[0] = 100
	firstArray[1] = 200
	firstArray[2] = 300

	// Get last position index
	// fmt.Println(len(firstArray) - 1)

	// Get last position value
	// fmt.Println(firstArray[len(firstArray)-1])

	for index, value := range firstArray {
		fmt.Printf("O meu indice é %d e o valor é %d\n", index, value)
	}
}
