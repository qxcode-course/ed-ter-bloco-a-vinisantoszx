package main

import (
	"bufio"
	"fmt"
	"os"
)

func busca(grid [][]byte, word string, i int, j int, pos int) bool {
	if pos == len(word) {
		return true
	}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) || grid[i][j] != word[pos] {
		return false
	}

	temp := grid[i][j]
	grid[i][j] = '#'

	achou := busca(grid, word, i+1, j, pos+1) ||
		busca(grid, word, i-1, j, pos+1) ||
		busca(grid, word, i, j+1, pos+1) ||
		busca(grid, word, i, j-1, pos+1)

	grid[i][j] = temp

	return achou
}

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == word[0] {
				if busca(grid, word, i, j, 0) {
					return true
				}
			}
		}
	}

	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
