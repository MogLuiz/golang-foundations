package main

import "fmt"

func main() {
	var myVar interface{} = "var name"
	println(myVar.(string))

	response, isConvertionOk := myVar.(int)
	fmt.Printf("O valor da resposta é %v e o resultado da tentativa de conversão é %v\n", response, isConvertionOk)
}
