package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data     []int
	size     int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand() {
	newCap := 1
	if ms.capacity > 0 {
		newCap = ms.capacity * 2
	}
	newData := make([]int, newCap)
	for i := 0; i < ms.size; i++ {
		newData[i] = ms.data[i]
	}
	ms.data = newData
	ms.capacity = newCap
}

func (ms *MultiSet) search(value int) (bool, int) {
	low := 0
	high := ms.size

	for low < high {
		mid := low + (high-low)/2

		if ms.data[mid] == value {
			return true, mid
		} else if ms.data[mid] < value {
			low = mid + 1
		} else {
			high = mid
		}
	}

	return false, low
}

func (ms *MultiSet) insert(value int, index int) error {
	if index < 0 || index > ms.size {
		return errors.New("index out of range")
	}
	if ms.size == ms.capacity {
		ms.expand()
	}
	for i := ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
	return nil
}

func (ms *MultiSet) Insert(value int) {
	_, index := ms.search(value)
	_ = ms.insert(value, index)
}

func (ms *MultiSet) Contains(value int) bool {
	found, _ := ms.search(value)
	return found
}

func (ms *MultiSet) erase(index int) error {
	if index < 0 || index >= ms.size {
		return errors.New("index out of range")
	}
	for i := index; i < ms.size-1; i++ {
		ms.data[i] = ms.data[i+1]
	}
	ms.size--
	return nil
}

func (ms *MultiSet) Erase(value int) bool {
	found, index := ms.search(value)
	if !found {
		return false
	}

	_ = ms.erase(index)
	return true
}

func (ms *MultiSet) String() string {
	if ms.size == 0 {
		return "[]"
	}
	var result strings.Builder
	result.WriteString("[")
	result.WriteString(fmt.Sprintf("%d", ms.data[0]))
	for i := 1; i < ms.size; i++ {
		result.WriteString(", ")
		result.WriteString(fmt.Sprintf("%d", ms.data[i]))
	}
	result.WriteString("]")
	return result.String()
}

func (ms *MultiSet) Count(value int) int {
	found, index := ms.search(value)
	if !found {
		return 0
	}

	count := 1

	for i := index - 1; i >= 0 && ms.data[i] == value; i-- {
		count++
	}
	for i := index + 1; i < ms.size && ms.data[i] == value; i++ {
		count++
	}

	return count
}

func (ms *MultiSet) Unique() int {
	if ms.size == 0 {
		return 0
	}

	uniqueCount := 1
	for i := 1; i < ms.size; i++ {
		if ms.data[i] != ms.data[i-1] {
			uniqueCount++
		}
	}
	return uniqueCount
}

func (ms *MultiSet) Clear() {
	ms.size = 0
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if !ms.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			if ms.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "count":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Count(value))
		case "unique":
			fmt.Println(ms.Unique())
		case "clear":
			ms.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
