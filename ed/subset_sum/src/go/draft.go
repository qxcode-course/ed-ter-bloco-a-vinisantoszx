package main

import "fmt"

func somaSub(subconj []int, k int, i int, somaAtual int) bool {
	if somaAtual == k {
		return true
	}

	if somaAtual > k || i == len(subconj) {
		return false
	}

    if somaSub(subconj, k, i + 1, somaAtual + subconj[i]) {
		return true
	}

    if somaSub(subconj, k, i + 1, somaAtual) {
		return true
	}

    return false
}

func main() {
	var n, k int

	fmt.Scan(&n, &k)

	subconj := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&subconj[i])
	}

    if somaSub(subconj, k, 0, 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
