package main
import "fmt"

type Pos struct {
    l, c int
}

func inside(grid [][]int, pos Pos) bool {
    nlinhas := len(grid)
    ncolunas := len(grid[0])

    return pos.l >= 0 && pos.l < nlinhas && pos.c >= 0 && pos.c < ncolunas
}

func (pos Pos) getNeig() []Pos {
    return []Pos {
        {pos.l - 1, pos.c},
        {pos.l + 1, pos.c},
        {pos.l, pos.c - 1},
        {pos.l, pos.c + 1},
    }
}

func ehFresca(grid [][]int, pos Pos) bool {
    return inside(grid, pos) && grid[pos.l][pos.c] == 1
}

func prepEstadoInicial(grid [][]int) ([]Pos, int) {
    fila := []Pos{}
    frescas := 0

    for i := 0; i < len(grid); i++ {
        for j := 0; j < len(grid[0]); j++ {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i, j})
            } else if grid[i][j] == 1 {
                frescas++
            }
        }
    }

    return fila, frescas
}

func orangesRotting(grid [][]int) int {
    fila, frescas := prepEstadoInicial(grid)

    if frescas == 0 {
        return 0
    }

    minutos := 0

    for len(fila) > 0 && frescas > 0 {
        tamanho := len(fila)

        for i := 0; i < tamanho; i++ {
            atual := fila[0]
            fila = fila[1:]

            for _, prox := range atual.getNeig() {
                if !ehFresca(grid, prox) {
                    continue
                }

                grid[prox.l][prox.c] = 2
                frescas--
                fila = append(fila, prox)
            }
        }

        minutos++
    }

    if frescas > 0 {
        return -1
    }

    return minutos
}

func main() {
    var linhas, colunas int

    fmt.Scan(&linhas, &colunas)

    grid := make([][]int, linhas)

    for i := 0; i < linhas; i++ {
        grid[i] = make([]int, colunas)

        for j := 0; j < colunas; j++ {
            fmt.Scan(&grid[i][j])
        }
    }

    fmt.Println(orangesRotting(grid))
}