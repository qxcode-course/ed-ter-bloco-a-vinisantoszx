package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b, c, p, heron float64

	fmt.Scan(&a, &b, &c)

	p = float64((a + b + c) / 2)

	heron = math.Sqrt(p * (p - a) * (p - b) * (p - c))

	fmt.Printf("%.2f\n", heron)
}
