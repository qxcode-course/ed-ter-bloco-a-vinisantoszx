package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		textoEntrada := scanner.Text()

		if textoEntrada == "" {
			continue
		}

		l := list.New()

		cursor := l.PushBack('|')

		for _, char := range textoEntrada {
			switch char {
			case 'R':
				l.InsertBefore('\n', cursor)
			case 'B':
				if vizinhoDeTras := cursor.Prev(); vizinhoDeTras != nil {
					l.Remove(vizinhoDeTras)
				}
			case 'D':
				if vizinhoDaFrente := cursor.Next(); vizinhoDaFrente != nil {
					l.Remove(vizinhoDaFrente)
				}
			case '<':
				if vizinhoDeTras := cursor.Prev(); vizinhoDeTras != nil {
					l.MoveBefore(cursor, vizinhoDeTras)
				}
			case '>':
				if vizinhoDaFrente := cursor.Next(); vizinhoDaFrente != nil {
					l.MoveAfter(cursor, vizinhoDaFrente)
				}
			default:
				l.InsertBefore(char, cursor)
			}
		}

		for node := l.Front(); node != nil; node = node.Next() {
			fmt.Printf("%c", node.Value.(rune))
		}
		fmt.Println()
	}
}
