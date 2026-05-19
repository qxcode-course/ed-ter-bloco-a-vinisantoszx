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
}

type List struct {
	root *Node
	size int
}

func NewList() *List {
	sentinel := &Node{}
	sentinel.next = sentinel
	sentinel.prev = sentinel
	return &List{
		root: sentinel,
		size: 0,
	}
}

func (ll *List) String() string {
	if ll.root.next == ll.root {
		return "[]"
	}

	var sb strings.Builder
	sb.WriteString("[")

	node := ll.root.next
	sb.WriteString(strconv.Itoa(node.Value))
	node = node.next

	for node != ll.root {
		sb.WriteString(", ")
		sb.WriteString(strconv.Itoa(node.Value))
		node = node.next
	}
	sb.WriteString("]")

	return sb.String()
}

func (ll *List) PushFront(value int) {
	newNode := &Node{Value: value}

	primeiroAtual := ll.root.next

	newNode.next = primeiroAtual
	newNode.prev = ll.root

	ll.root.next = newNode
	primeiroAtual.prev = newNode

	ll.size++
}

func (ll *List) PushBack(value int) {
	newNode := &Node{Value: value}

	ultimoAtual := ll.root.prev

	newNode.prev = ultimoAtual
	newNode.next = ll.root

	ultimoAtual.next = newNode
	ll.root.prev = newNode

	ll.size++
}

func (ll *List) Size() int {
	return ll.size
}

func (ll *List) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
	ll.size = 0
}

func (ll *List) PopBack() {
	if ll.size == 0 {
		return
	}

	alvo := ll.root.prev
	vizinhoDeTras := alvo.prev

	ll.root.prev = vizinhoDeTras
	vizinhoDeTras.next = ll.root

	ll.size--
}

func (ll *List) PopFront() {
	if ll.size == 0 {
		return
	}

	alvo := ll.root.next
	vizinhoDaFrente := alvo.next

	ll.root.next = vizinhoDaFrente
	vizinhoDaFrente.prev = ll.root

	ll.size--
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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
