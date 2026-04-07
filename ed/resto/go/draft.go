package main

import "fmt"

func div(n int) {
	if n == 0 {
		return
	}

	resultado := n / 2
	resto := n % 2

	div(resultado)

	fmt.Printf("%d %d\n", resultado, resto)
}

func main() {
	var n int

	fmt.Scan(&n)

	div(n)
}
