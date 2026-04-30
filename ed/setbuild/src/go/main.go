package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Set struct {
	data     []int
	size     int
	capacity int
}

func NewSet(capacity int) *Set {
	return &Set{
		data:     make([]int, capacity),
		size:     0,
		capacity: capacity,
	}
}

func (s *Set) reserve(newCapacity int) {
	if newCapacity > s.capacity {
		newData := make([]int, newCapacity)
		for i := 0; i < s.size; i++ {
			newData[i] = s.data[i]
		}
		s.data = newData
		s.capacity = newCapacity
	}
}

func (s *Set) binarySearch(value int) int {
	left := 0
	right := s.size - 1

	for left <= right {
		mid := left + (right-left)/2
		if s.data[mid] == value {
			return mid
		}
		if s.data[mid] < value {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func (s *Set) insert(value int, index int) error {
	if index < 0 || index > s.size {
		return errors.New("index out of range")
	}

	if s.size == s.capacity {
		newCap := 1
		if s.capacity > 0 {
			newCap = s.capacity * 2
		}
		s.reserve(newCap)
	}

	for i := s.size; i > index; i-- {
		s.data[i] = s.data[i-1]
	}

	s.data[index] = value
	s.size++
	return nil
}

func (s *Set) Insert(value int) {
	if s.binarySearch(value) != -1 {
		return
	}

	insertIndex := 0
	for insertIndex < s.size && s.data[insertIndex] < value {
		insertIndex++
	}

	_ = s.insert(value, insertIndex)
}

func (s *Set) Contains(value int) bool {
	return s.binarySearch(value) != -1
}

func (s *Set) erase(index int) error {
	if index < 0 || index >= s.size {
		return errors.New("index out of range")
	}

	for i := index; i < s.size-1; i++ {
		s.data[i] = s.data[i+1]
	}
	s.size--
	return nil
}

func (s *Set) Erase(value int) bool {
	index := s.binarySearch(value)
	if index == -1 {
		return false
	}

	_ = s.erase(index)
	return true
}

func (s *Set) Clear() {
	s.size = 0
}

func (s *Set) String() string {
	if s.size == 0 {
		return "[]"
	}
	var result strings.Builder
	result.WriteString("[")
	result.WriteString(fmt.Sprintf("%d", s.data[0]))
	for i := 1; i < s.size; i++ {
		result.WriteString(", ")
		result.WriteString(fmt.Sprintf("%d", s.data[i]))
	}
	result.WriteString("]")
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v)
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
