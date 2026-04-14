# Rascunho

Se a tarefa exigir um relatório, escreva ele aqui. Você pode usar markdown, imagens e o que mais quiser para criar um relatório bem completo.
package main

import (
	"fmt"
)

func circulos(pen *Pen, raio float64, nivel int) {
	if raio < 10 {
		return
	}

	pen.SetRGB(255, 255, 255)

	pen.DrawCircle(raio)

	for range 6{
		pen.Up()
		pen.Walk(raio)
		pen.Down()

		circulos(pen, raio*0.40, nivel-1)

		pen.Right(90)
	
		
		pen.Left(90)
		pen.Up()
		pen.Walk(-raio)
		pen.Down()
	}

}

func main() {
	pen := NewPen(1200, 1200)
	pen.SetPosition(600, 600)
	circulos(pen, 320, 5)
	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}
