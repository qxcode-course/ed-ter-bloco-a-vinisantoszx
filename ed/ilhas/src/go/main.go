package main

import (
	"bufio"
	"fmt"
	"os"
)

func afundarIlha(grid [][]byte, l, c int) {
	linhas := len(grid)
	colunas := len(grid[0])

	if l < 0 || l >= linhas || c < 0 || c >= colunas || grid[l][c] == '0' {
		return
	}

	grid[l][c] = '0'

	afundarIlha(grid, l-1, c)
	afundarIlha(grid, l+1, c)
	afundarIlha(grid, l, c-1)
	afundarIlha(grid, l, c+1)
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	nIlha := 0

	linhas := len(grid)
	colunas := len(grid[0])

	for l := 0; l < linhas; l++ {
		for c := 0; c < colunas; c++ {
			if grid[l][c] == '1' {
				nIlha++

				afundarIlha(grid, l, c)
			}
		}
	}

	return nIlha
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
