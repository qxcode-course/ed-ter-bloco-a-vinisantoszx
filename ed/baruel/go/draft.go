package main

import "fmt"

func main() {
	var albumsize, possuisize int

	fmt.Scan(&albumsize, &possuisize)

	figuras := make([]int, possuisize)

	for i := range figuras {
		fmt.Scanf("%d", &figuras[i])
	}

	unicos := make(map[int]bool)
	repetidos := make([]int, 0, possuisize)

	for _, figura := range figuras {
		if unicos[figura] {
			repetidos = append(repetidos, figura)
		} else {
			unicos[figura] = true
		}
	}

	faltantes := make([]int, 0, albumsize)

    for i := 1; i <= albumsize; i++ {
        if !unicos[i] {
            faltantes = append(faltantes, i)
        }
    }

	saida := fmt.Sprintf("%v", repetidos)
    if saida == "[]" {
        fmt.Println("N")
    } else {
        fmt.Println(saida[1 : len(saida)-1])
    }

	saida = fmt.Sprintf("%v", faltantes)
    if saida == "[]" {
        fmt.Println("N")
    } else {
        fmt.Println(saida[1 : len(saida)-1])
    }

}
