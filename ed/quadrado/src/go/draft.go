package main
import "fmt"

func quadrado(n int) int {
    if n == 1 {
        fmt.Println("1^2 = 1")
        return 1
    }

    anterior := n - 1

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = ?\n", n, anterior, anterior)

    resultado := quadrado(anterior) + 2*anterior + 1

    fmt.Printf("%d^2 = %d^2 + 2*%d + 1 = %d\n", n, anterior, anterior, resultado)

    return resultado
}

func main() {
    var n int

    fmt.Scan(&n)
    
    quadrado(n)
}