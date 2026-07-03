package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Point struct {
	row int
	col int
}

func readMaze(reader *bufio.Reader, rows int, cols int) ([][]rune, Point, Point) {
	maze := make([][]rune, rows)

	start := Point{}
	end := Point{}

	for i := 0; i < rows; i++ {
		line, _ := reader.ReadString('\n')
		line = strings.TrimRight(line, "\r\n")

		maze[i] = []rune(line)

		for j := 0; j < cols; j++ {
			if maze[i][j] == 'I' {
				start = Point{i, j}
			}

			if maze[i][j] == 'F' {
				end = Point{i, j}
			}
		}
	}

	return maze, start, end
}

func isValid(point Point, maze [][]rune, visited [][]bool, rows int, cols int) bool {
	if point.row < 0 || point.row >= rows || point.col < 0 || point.col >= cols ||
		maze[point.row][point.col] == '#' || visited[point.row][point.col] {
		return false
	}

	return true
}

func resolveMaze(maze [][]rune, start Point, end Point, rows int, cols int) {
	caminho := NewStack[Point]()

	visited := make([][]bool, rows)

	for i := 0; i < rows; i++ {
		visited[i] = make([]bool, cols)
	}

	caminho.Push(start)

	directions := []Point{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}

	for !caminho.IsEmpty() {
		current := caminho.Top()
		visited[current.row][current.col] = true

		if current == end {
			break
		}

		foundNext := false

		for _, direction := range directions {
			next := Point{
				row: current.row + direction.row,
				col: current.col + direction.col,
			}

			if isValid(next, maze, visited, rows, cols) {
				caminho.Push(next)
				foundNext = true
				break
			}
		}

		if !foundNext {
			caminho.Pop()
		}
	}

	for !caminho.IsEmpty() {
		point := caminho.Pop()
		maze[point.row][point.col] = '.'
	}
}

func printMaze(maze [][]rune) {
	for i := 0; i < len(maze); i++ {
		fmt.Println(string(maze[i]))
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	var rows, cols int
	fmt.Fscan(reader, &rows, &cols)

	reader.ReadString('\n')

	maze, start, end := readMaze(reader, rows, cols)

	resolveMaze(maze, start, end, rows, cols)

	printMaze(maze)
}
