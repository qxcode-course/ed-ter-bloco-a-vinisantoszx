package main

import (
	"fmt"
)

func resolver(textoEdit []byte, l int, i int) bool {
	if i == len(textoEdit) {
		return true
	}

	if textoEdit[i] != '.' {
		return resolver(textoEdit, l, i+1)
	}

	for num := 0; num <= l; num++ {
		charAtual := byte('0' + num)

		if colocar(textoEdit, l, i, charAtual) {
			textoEdit[i] = charAtual

			if resolver(textoEdit, l, i+1) {
				return true
			}

			textoEdit[i] = '.'
		}
	}

	return false
}

func colocar(textoEdit []byte, l int, i int, charAtual byte) bool {
	for passo := 1; passo <= l; passo++ {
		esq := i - passo

		if esq >= 0 && textoEdit[esq] == charAtual {
			return false
		}

		dir := i + passo

		if dir < len(textoEdit) && textoEdit[dir] == charAtual {
			return false
		}
	}
	return true
}

func main() {
	var sequencia string
	var limite int

	fmt.Scan(&sequencia)
	fmt.Scan(&limite)

	textoEdit := []byte(sequencia)

	resolver(textoEdit, limite, 0)

	fmt.Println(string(textoEdit))
}
