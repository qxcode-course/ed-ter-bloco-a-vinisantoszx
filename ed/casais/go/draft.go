package main

import (
	"fmt"
)
func main() {
    var vsize, ncasais int
    fmt.Scan(&vsize)

    casais := make([]int, vsize)

    for i := 0; i < vsize; i++ {
        fmt.Scan(&casais[i])
    }

    for i := 0; i < vsize; i++ {
        for j := 0; j < vsize; j++ {
            if casais[i] > 0 {
                if casais[i] == (casais[j] * (-1)) {
                    casais[j] = 0
                    ncasais++
                }
            }
        }
    }

    fmt.Println(ncasais)
}
