package main

import (
	"fmt"
)

func circulos(pen *Pen, raio float64, nivel int) {
	if nivel == 0 {
		return
	}

	pen.SetRGB(255, 255, 255)

	pen.DrawCircle(raio)

	for i := 0; i < 6; i++ {
		pen.Right(60)

		pen.Up()
		pen.Walk(raio)
		pen.Down()

		pen.Right(90)
		
		circulos(pen, raio*0.35, nivel-1)
		
		pen.Left(90)
		pen.Up()
		pen.Walk(-raio)
		pen.Down()
	}

}

func main() {
	pen := NewPen(1200, 1200)
	pen.SetPosition(600, 600)
	circulos(pen, 300, 5)
	pen.SavePNG("circulos.png")
	fmt.Println("PNG file created successfully.")
}
