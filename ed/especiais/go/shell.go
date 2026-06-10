package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Pair struct {
	One int
	Two int
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func occurr(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}

	counts := make(map[int]int)
	for _, v := range vet {
		counts[abs(v)]++
	}

	var keys []int
	for k := range counts {
		keys = append(keys, k)
	}

	sort.Ints(keys)

	var result []Pair
	for _, k := range keys {
		result = append(result, Pair{One: k, Two: counts[k]})
	}

	return result
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{}
	}

	var result []Pair

	currentVal := abs(vet[0])
	count := 1

	for i := 1; i < len(vet); i++ {
		val := abs(vet[i])

		if val == currentVal {
			count++
		} else {
			result = append(result, Pair{One: currentVal, Two: count})
			currentVal = val
			count = 1
		}
	}

	result = append(result, Pair{One: currentVal, Two: count})
	
	return result
}

func mnext(vet []int) []int {
	result := make([]int, len(vet))

	if len(vet) == 0 {
		return result
	}

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			temMulherAoLado := false

			if i > 0 && vet[i-1] < 0 {
				temMulherAoLado = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				temMulherAoLado = true
			}

			if temMulherAoLado {
				result[i] = 1
			}
		}
	}

	return result
}

func alone(vet []int) []int {
	result := make([]int, len(vet))

	if len(vet) == 0 {
		return result
	}

	for i := 0; i < len(vet); i++ {
		if vet[i] > 0 {
			temMulherAoLado := false

			if i > 0 && vet[i-1] < 0 {
				temMulherAoLado = true
			}

			if i < len(vet)-1 && vet[i+1] < 0 {
				temMulherAoLado = true
			}

			if !temMulherAoLado {
				result[i] = 1
			}
		}
	}

	return result
}

func couple(vet []int) int {
	couplesCount := 0

	disponivel := make([]bool, len(vet))
	for i := range disponivel {
		disponivel[i] = true
	}

	for i := 0; i < len(vet); i++ {
		if !disponivel[i] {
			continue
		}

		for j := i + 1; j < len(vet); j++ {
			if !disponivel[j] {
				continue
			}

			if vet[i]+vet[j] == 0 {
				couplesCount++
				disponivel[i] = false
				disponivel[j] = false
				break
			}
		}
	}

	return couplesCount
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	for i := 0; i < len(seq); i++ {
		if vet[pos+i] != seq[i] {
			return false
		}
	}

	return true
}

func subseq(vet []int, seq []int) int {
	if len(seq) == 0 || len(seq) > len(vet) {
		return -1
	}

	limite := len(vet) - len(seq)

	for i := 0; i <= limite; i++ {
		if hasSubseq(vet, seq, i) {
			return i
		}
	}

	return -1
}

func erase(vet []int, posList []int) []int {
	toDelete := make(map[int]bool)
	for _, pos := range posList {
		toDelete[pos] = true
	}

	var result []int

	for i, v := range vet {
		if !toDelete[i] {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return []int{}
	}

	return result
}

func clear(vet []int, value int) []int {
	var result []int

	for _, v := range vet {
		if v != value {
			result = append(result, v)
		}
	}

	if len(result) == 0 {
		return []int{}
	}

	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
