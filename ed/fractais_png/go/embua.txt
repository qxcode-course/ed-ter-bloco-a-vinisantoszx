package main

import (
	"fmt"
	"math/rand"
)

func ri(inf, sup int) float64 {
	return float64(rand.Intn(sup-inf+1) + inf)
}

func arvere(pen *Pen, dist float64) {
	if dist < 10 {
		if ri(0, 50) == 0 {
			pen.SetRGB(255, 0, 0)
			pen.FillCircle(10)
		}
		return
	}
	ang_dir := ri(10, 40)
	ang_esq := ri(10, 40)

	pen.SetLineWidth(dist / 5)
	pen.SetRGB(0, 0, 0)
	pen.Walk(dist)
	pen.Right(ang_dir)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Left(ang_dir + ang_esq)
	arvere(pen, dist*(ri(80, 85)/100))
	pen.Right(ang_esq)
	pen.SetRGB(0, 0, 0)
	pen.Walk(-dist)
}

func main() {
	pen := NewPen(600, 500)
	pen.SetHeading(90)
	pen.SetPosition(300, 500)
	arvere(pen, 70)
	pen.SavePNG("tree.png")
	fmt.Println("PNG file created successfully.")
}
