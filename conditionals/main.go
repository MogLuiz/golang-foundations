package main

func main() {
	a := 1
	b := 2
	c := 3

	if a > b && c > a {
		println(a)
	} else if b > a {
		println(b)
	}

	switch c {
	case 1:
		println("a")
	case 2:
		println("b")
	default:
		println("x")
	}
}
