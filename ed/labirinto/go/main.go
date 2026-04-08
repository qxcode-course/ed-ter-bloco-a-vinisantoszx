package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func search(grid [][]rune, curr Pos, end Pos) bool {
	if curr.l < 0 || curr.l >= len(grid) || curr.c < 0 || curr.c >= len(grid[0]) || grid[curr.l][curr.c] != ' ' {
		return false
	}

	grid[curr.l][curr.c] = '.'

	if curr == end {
		return true
	}

	vizinhos := []Pos{
		{curr.l - 1, curr.c},
		{curr.l + 1, curr.c},
		{curr.l, curr.c - 1},
		{curr.l, curr.c + 1},
	}

	for _, vizinho := range vizinhos {
		if search(grid, vizinho, end) {
			return true
		}
	}

	grid[curr.l][curr.c] = ' '
	return false
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nl_nc := scanner.Text()
	var nl, nc int
	fmt.Sscanf(nl_nc, "%d %d", &nl, &nc)
	grid := make([][]rune, nl)

	// Lê a gridriz
	for i := range nl {
		scanner.Scan()
		grid[i] = []rune(scanner.Text())
	}

	// Procura posições de início e endPos e conserta para _
	var startPos, endPos Pos
	for l := range nl {
		for c := range nc {
			if grid[l][c] == 'I' {
				grid[l][c] = ' '
				startPos = Pos{l, c}
			}
			if grid[l][c] == 'F' {
				grid[l][c] = ' '
				endPos = Pos{l, c}
			}
		}
	}

	search(grid, startPos, endPos)

	// Imprime o labirinto final
	for _, line := range grid {
		fmt.Println(string(line))
	}
}
