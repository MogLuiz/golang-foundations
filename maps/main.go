package main

import "fmt"

func main() {
	salary := map[string]int{"Luiz": 10000, "Lucas": 12000}
	fmt.Println(salary["Luiz"])

	delete(salary, "Luiz")
	fmt.Println(salary)

	salary["Bruno"] = 4000
	fmt.Println(salary)

	// Criando maps sem inicializar com valores
	withoutValues := make(map[string]int)
	withoutValues2 := map[string]int{}
	fmt.Println(withoutValues, withoutValues2)

	withoutValues["Gislene"] = 2000

	// Percorrendo maps
	team := map[string]string{"Luiz": "Engineer", "Lucas": "Manager", "Julio": "Tech Lead"}

	for key, value := range team {
		fmt.Printf("O colaborador %s tem o seguinte cargo: %s\n", key, value)
	}

	// Ignorando a KEY do map
	for _, value := range team {
		fmt.Printf("O colaborador tem o seguinte cargo: %s\n", value)
	}
}
