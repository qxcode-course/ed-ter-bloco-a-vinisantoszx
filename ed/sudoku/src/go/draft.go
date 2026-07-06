package main

import (
	"bufio"
	"fmt"
	"os"
)

func existeNaLinha(grid [][]rune, linha int, value rune) bool {
    n := len(grid)

    for coluna := 0; coluna < n; coluna++ {
        if grid[linha][coluna] == value {
            return true
        }
    }

    return false
}

func existeNaColuna(grid [][]rune, coluna int, value rune) bool {
    n := len(grid)

    for linha := 0; linha < n; linha++ {
        if grid[linha][coluna] == value {
            return true
        }
    }

    return false
}

func tamanhoQuadrante(n int) int {
    if n == 4 {
        return 2
    }

    return 3
}

func existeNoQuadrante(grid [][]rune, linha int, coluna int, value rune) bool {
    n := len(grid)
    tam := tamanhoQuadrante(n)

    linhaInicio := (linha / tam) * tam
    colunaInicio := (coluna / tam) * tam

    for i := linhaInicio; i < linhaInicio+tam; i++ {
        for j := colunaInicio; j < colunaInicio+tam; j++ {
            if grid[i][j] == value {
                return true
            }
        }
    }

    return false
}

func podeColocar(grid [][]rune, linha int, coluna int, value rune) bool {
    if existeNaLinha(grid, linha, value) {
        return false
    }

    if existeNaColuna(grid, coluna, value) {
        return false
    }

    if existeNoQuadrante(grid, linha, coluna, value) {
        return false
    }

    return true
}

func resolver(grid [][]rune, i int) bool{
    n := len(grid)

    if i == n*n {
        return true
    }

    linha := i / n
    coluna := i % n

    if grid[linha][coluna] != '.' {
        return resolver(grid, i+1)
    }

    for value := '1'; value <= rune('0'+n); value++ {
        if podeColocar(grid, linha, coluna, value) {
            grid[linha][coluna] = value

            if resolver(grid, i+1) {
                return true
            }

            grid[linha][coluna] = '.'
        }
    }

    return false
}

func imprimir(grid [][]rune) {
    for i := 0; i < len(grid); i++ {
        fmt.Println(string(grid[i]))
    }
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)

    var n int
    fmt.Scan(&n)

    grid := make([][]rune, n)

    for i := 0; i < n; i++ {
        scanner.Scan()
        linha := scanner.Text()
        grid[i] = []rune(linha)
    }

    resolver(grid, 0)
    imprimir(grid)
}
