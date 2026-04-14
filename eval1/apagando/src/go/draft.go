package main

import "fmt"

func main() {
	var nfila, nsaiu int

	fmt.Scan(&nfila)

	fila := make([]int, nfila)

	for i := 0; i < nfila; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&nsaiu)

	saiu := make([]int, nsaiu)

	for i := 0; i < nsaiu; i++ {
		fmt.Scan(&saiu[i])
	}

	sobrou := make(map[int]bool)

	for i := 0; i < nsaiu; i++ {
		sobrou[saiu[i]] = true
	}

	for i := 0; i < nfila; i++ {
		if !sobrou[fila[i]] {
			fmt.Print(fila[i])
			fmt.Print(" ")
		}
	}

	fmt.Println()
}
