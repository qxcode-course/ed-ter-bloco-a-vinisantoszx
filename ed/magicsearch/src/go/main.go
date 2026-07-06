package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func MagicSearch(slice []int, value int) int {
	left := 0
	right := len(slice) - 1

	for left <= right {
		mid := (left + right) / 2

		if slice[mid] == value {
			for mid+1 < len(slice) && slice[mid+1] == value {
				mid++
			}

			return mid
		}

		if slice[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
