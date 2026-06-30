package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	nl := len(grid)
	nc := len(grid[0])

	stack.Push(Pos{l: l, c: c})

	for !stack.IsEmpty() {
		atual := stack.Pop()

		if atual.l < 0 || atual.l >= nl || atual.c < 0 || atual.c >= nc {
			continue
		}

		if grid[atual.l][atual.c] == '#' {
			grid[atual.l][atual.c] = 'o'

			stack.Push(Pos{l: atual.l - 1, c: atual.c})
			stack.Push(Pos{l: atual.l + 1, c: atual.c})
			stack.Push(Pos{l: atual.l, c: atual.c - 1})
			stack.Push(Pos{l: atual.l, c: atual.c + 1})
		}
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
