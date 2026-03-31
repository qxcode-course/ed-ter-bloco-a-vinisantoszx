package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func getMen(vet []int) []int {
	var homem []int

	for _, pessoa := range vet {
		if pessoa > 0 {
			homem = append(homem, pessoa)
		}
	}

	return homem
}

func getCalmWomen(vet []int) []int {
	var mulher []int

	for _, pessoa := range vet {
		if pessoa < 0 && pessoa > -10 {
			mulher = append(mulher, pessoa)
		}
	}

	return mulher
}

func sortVet(vet []int) []int {
	var ordem []int

	ordem = append(ordem, vet...)

	sort.Ints(ordem)

	return ordem
}

func sortStress(vet []int) []int {
	var ordem []int

	ordem = append(ordem, vet...)

	sort.Slice(ordem, func(i, j int) bool {
		return math.Abs(float64(ordem[i])) < math.Abs(float64(ordem[j]))
	})

	return ordem
}

func reverse(vet []int) []int {
	var reverso []int

	for i := len(vet) - 1; i >= 0; i-- {
		reverso = append(reverso, vet[i])
	}

	return reverso
}

func unique(vet []int) []int {
	var unico []int

	visto := make(map[int]bool)

	for _, p := range vet {
		if !visto[p] {
			unico = append(unico, p)
		}
		visto[p] = true
	}

	return unico
}

func repeated(vet []int) []int {
	var repetido []int

	visto := make(map[int]bool)

	for _, p := range vet {
		if visto[p] {
			repetido = append(repetido, p)
		}
		visto[p] = true
	}

	return repetido
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
