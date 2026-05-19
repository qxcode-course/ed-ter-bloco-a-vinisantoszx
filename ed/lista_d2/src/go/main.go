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
	next  *Node
	prev  *Node
	root  *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

type List struct {
	root *Node
	size int
}

func NewList() *List {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	sentinel.root = sentinel
	return &List{
		root: sentinel,
		size: 0,
	}
}

func (ll *List) Front() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.next
}

func (ll *List) Back() *Node {
	if ll.size == 0 {
		return nil
	}
	return ll.root.prev
}

func (ll *List) PushFront(value int) {
	newNode := &Node{Value: value, root: ll.root}
	primeiroAtual := ll.root.next

	newNode.next = primeiroAtual
	newNode.prev = ll.root

	ll.root.next = newNode
	primeiroAtual.prev = newNode
	ll.size++
}

func (ll *List) PushBack(value int) {
	newNode := &Node{Value: value, root: ll.root}
	ultimoAtual := ll.root.prev

	newNode.prev = ultimoAtual
	newNode.next = ll.root

	ultimoAtual.next = newNode
	ll.root.prev = newNode
	ll.size++
}

func (ll *List) String() string {
	if ll.size == 0 {
		return "[]"
	}
	var sb strings.Builder
	sb.WriteString("[")

	node := ll.Front()
	sb.WriteString(strconv.Itoa(node.Value))
	node = node.Next()

	for node != nil {
		sb.WriteString(", ")
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.Next()
	}
	sb.WriteString("]")
	return sb.String()
}

func (ll *List) PopFront() {
	if primeiro := ll.Front(); primeiro != nil {
		ll.Remove(primeiro)
	}
}

func (ll *List) PopBack() {
	if ultimo := ll.Back(); ultimo != nil {
		ll.Remove(ultimo)
	}
}

func (ll *List) Search(value int) *Node {
	for node := ll.Front(); node != nil; node = node.Next() {
		if node.Value == value {
			return node
		}
	}
	return nil
}

func (ll *List) Insert(node *Node, value int) {
	newNode := &Node{Value: value, root: ll.root}

	vizinhoDeTras := node.prev

	newNode.prev = vizinhoDeTras
	newNode.next = node

	vizinhoDeTras.next = newNode
	node.prev = newNode

	ll.size++
}

func (ll *List) Remove(node *Node) *Node {
	if node == nil || node == ll.root {
		return nil
	}

	nextNode := node.Next() 
	
	vizinhoDeTras := node.prev
	vizinhoDaFrente := node.next

	vizinhoDeTras.next = vizinhoDaFrente
	vizinhoDaFrente.prev = vizinhoDeTras
	
	ll.size--
	
	return nextNode
}

func (ll *List) Size() int {
	return ll.size
}

func (ll *List) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
