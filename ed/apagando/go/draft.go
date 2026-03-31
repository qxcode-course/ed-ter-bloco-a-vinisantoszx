package main

import "fmt"

// ANOTAÇÃO PESSOAL - O SCANF NÃO FUNCIONOU NESSE CASO, TIRAR DÚVIDA COM O PROFESSOR NA AULA DE ED
func main() {
	var n, m, id int

	fmt.Scan(&n)

	fila := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&fila[i])
	}

	fmt.Scan(&m)

	desistiu := make(map[int]bool)

	for i := 0; i < m; i++ {
		fmt.Scan(&id)
		desistiu[id] = true
	}

	for _, pessoa := range fila {
		if !desistiu[pessoa] {
			fmt.Printf("%d ", pessoa)
		}
	}

	fmt.Println()
}
