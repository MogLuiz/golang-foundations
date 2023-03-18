package main

import "fmt"

// O tipo de interface vazia implemente todo mundo. (Usar com sabedoria)

type emptyInterface interface{}

func main() {
	var x emptyInterface = 10
	var y emptyInterface = "Oi"

	showType(x)
	showType(y)
}

func showType(t interface{}) {
	fmt.Printf("O tipo da variável é %T e o valor é %v\n", t, t)
}
