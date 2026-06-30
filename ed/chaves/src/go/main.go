package main

import "fmt"

func main() {
	q := NewQueue[string]()

	for i := 0; i < 16; i++ {
		equipe := string(rune('A' + i))
		q.Enqueue(equipe)
	}

	for i := 0; i < 15; i++ {
		var g1, g2 int

		fmt.Scan(&g1, &g2)

		t1 := q.Dequeue()
		t2 := q.Dequeue()

		if g1 > g2 {
			q.Enqueue(t1)
		} else {
			q.Enqueue(t2)
		}
	}

	campeao := q.Dequeue()
	fmt.Println(campeao)
}
