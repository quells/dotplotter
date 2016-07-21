package main

import (
	dp "github.com/quells/dotplotter"
)

func main() {
	// dp.NewCanvas(w, h, xmin, xmax, ymin, ymax)
	c := dp.NewCanvas(600, 600, -10, 10, -10, 10)

	// c.DrawCircleAt(modelX, modelY, pixelRadius, color)
	c.DrawCircleAt(0, 0, 10, dp.DefaultColor("black"))
	c.DrawCircleAt(-1, 1, 10, dp.DefaultColor("red"))
	c.DrawCircleAt(0, 1, 10, dp.DefaultColor("orange"))
	c.DrawCircleAt(1, 1, 10, dp.DefaultColor("yellow"))
	c.DrawCircleAt(-1, -1, 10, dp.DefaultColor("green"))
	c.DrawCircleAt(0, -1, 10, dp.DefaultColor("blue"))
	c.DrawCircleAt(1, -1, 10, dp.DefaultColor("purple"))

	c.ExportToPNG("test.png")
}
