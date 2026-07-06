package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func longestIncreasingPath(matrix [][]int) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	rows := len(matrix)
	cols := len(matrix[0])

	salva := make([][]int, rows)

	for i := 0; i < rows; i++ {
		salva[i] = make([]int, cols)
	}

	directions := [][]int {
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	var dfs func(linha int, coluna int) int

	dfs = func (linha int, coluna int) int {
		if salva[linha][coluna] != 0 {
			return salva[linha][coluna]
		}

		melhor := 1

		for _, direction := range directions {
			proximaLinha := linha + direction[0]
			proximaColuna := coluna + direction[1]

			if proximaLinha < 0 || proximaLinha >= rows || proximaColuna < 0 || proximaColuna >= cols {
				continue
			}

			if matrix[proximaLinha][proximaColuna] <= matrix[linha][coluna] {
				continue
			}

			tamCaminho := 1 + dfs(proximaLinha, proximaColuna)

			if tamCaminho > melhor {
				melhor = tamCaminho
			}
		}

		salva[linha][coluna] = melhor
		return melhor
	}

	maiorCaminho := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			caminhoAtual := dfs(i, j)

			if caminhoAtual > maiorCaminho {
				maiorCaminho = caminhoAtual
			}
		}
	}

	return maiorCaminho
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
