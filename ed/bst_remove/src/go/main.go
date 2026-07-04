package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func BstInsert(values []int) *Node {
	if len(values) == 0 {
		return nil
	}

	raiz := &Node{Value: values[0]}

	for _, v := range values[1:] {
		atual := raiz

		for {
			if v < atual.Value {
				if atual.Left == nil {
					atual.Left = &Node{Value: v}
					break
				}

				atual = atual.Left
			} else if v > atual.Value {
				if atual.Right == nil {
					atual.Right = &Node{Value: v}
					break
				}

				atual = atual.Right
			} else {
				break
			}
		}
	}

	return raiz
}

func BstRemove(node *Node, value int) *Node {
	if node == nil {
		return nil
	}

	if value < node.Value {
		node.Left = BstRemove(node.Left, value)
	} else if value > node.Value {
		node.Right = BstRemove(node.Right, value)
	} else {
		if node.Left == nil && node.Right == nil {
			return nil
		}

		if node.Left == nil {
			return node.Right
		}

		if node.Right == nil {
			return node.Left
		}

		antecessor := node.Left

		for antecessor.Right != nil {
			antecessor = antecessor.Right
		}

		node.Value = antecessor.Value
		node.Left = BstRemove(node.Left, antecessor.Value)
	}

	return node
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(parts))
	for _, elem := range parts {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	scanner.Scan()
	toRemove, _ := strconv.Atoi(scanner.Text())

	_ = toRemove // Ignora o valor a ser removido, pois não está implementado
	root := BstInsert(values)
	fmt.Println("original:")
	BShow(root, "") // Chama a função de impressão formatada
	root = BstRemove(root, toRemove)
	fmt.Println("modificado:")
	BShow(root, "") // Chama a função de impressão formatada da árvore modificada
}
