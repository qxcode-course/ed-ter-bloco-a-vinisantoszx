package main

import "fmt"

// TRABALHO DMS PRA ENTENDER OS ESPAÇOS DESSES PRINTS KKKKKKKKKKK
func imprimirFila(v []int, espada int) {
	fmt.Print("[ ")
	for i, p := range v {
		if p != 0 {
			fmt.Print(p)
			if i == espada {
				fmt.Print("> ")
			} else {
				fmt.Print(" ")
			}
		}
	}
	fmt.Println("]")
}

func main() {
	var n, e int

	fmt.Scanf("%d %d", &n, &e)

	v := make([]int, n)
	for i := 0; i < n; i++ {
		v[i] = i + 1
	}

	espada := e - 1
	vivos := n

	imprimirFila(v, espada)

	for vivos > 1 {
		vitima := (espada + 1) % n
		for v[vitima] == 0 {
			vitima = (vitima + 1) % n
		}

		v[vitima] = 0
		vivos--

		espada = (vitima + 1) % n
		for v[espada] == 0 {
			espada = (espada + 1) % n
		}

		imprimirFila(v, espada)
	}
}
