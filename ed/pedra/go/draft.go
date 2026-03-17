package main

import (
	"fmt"
)

func main() {
	var comp, a, b, menor, diferenca, indice int

	fmt.Scanf("%d", &comp)

	menor = 999

	for i := 0; i < comp; i++{
		fmt.Scanf("%d %d", &a, &b)

		if a >= 10 && b >= 10 {
			diferenca = a - b

			if diferenca < 0 {
				diferenca = -diferenca
			}

			if menor > diferenca {
				menor = diferenca
				indice = i
			}
		}
	}

	if indice != 0 {
		fmt.Println(indice)
	} else {
		fmt.Println("sem ganhador")
	}

}
