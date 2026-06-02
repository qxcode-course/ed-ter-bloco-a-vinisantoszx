package main

import "fmt"

func main() {
	var q int
	var d string

	fmt.Scan(&q, &d)

	x := make([]int, q)
	y := make([]int, q)

	for i := 0; i < q; i++ {
		fmt.Scan(&x[i], &y[i])
	}

	for i := q - 1; i > 0; i-- {
		x[i] = x[i-1]
		y[i] = y[i-1]
	}

	switch d {
	case "L":
		x[0]--
	case "R":
		x[0]++
	case "U":
		y[0]--
	case "D":
		y[0]++
	default:
		fmt.Println("Digite uma opcao valida")
	}

	for i := 0; i < q; i++ {
		fmt.Println(x[i], y[i])
	}
}
