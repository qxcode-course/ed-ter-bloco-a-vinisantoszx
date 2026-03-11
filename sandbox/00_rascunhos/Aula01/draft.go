package main

import "fmt"

func main() {
	var idade int

	fmt.Scanf("%d", &idade)

	if idade > 0 {
		if idade < 16 {
			fmt.Println("Você não pode votar")
		} else if idade >= 16 && idade < 18 {
			fmt.Println("Você tem direito de votar")
		} else if idade >= 18 && idade < 65 {
			fmt.Println("Você é obrigado a votar")
		} else if idade >= 65 {
			fmt.Println("Você tem direito de votar, mas não é obrigado")
		}
	} else {
		fmt.Println("Digite um valor válido")
	}

}
