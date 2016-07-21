// dotplotter provides utilities for creating simple images from 2D coordinates.
package dotplotter

import (
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"path/filepath"
)

// DefaultColor exports default RGB values for common colors.
func DefaultColor(c string) color.RGBA {
	m := map[string]color.RGBA{
		"white":  color.RGBA{255, 255, 255, 255},
		"black":  color.RGBA{0, 0, 0, 255},
		"red":    color.RGBA{255, 0, 0, 255},
		"orange": color.RGBA{255, 128, 0, 255},
		"yellow": color.RGBA{255, 255, 0, 255},
		"green":  color.RGBA{0, 192, 0, 255},
		"blue":   color.RGBA{0, 0, 255, 255},
		"purple": color.RGBA{128, 0, 255, 255},
	}
	r, ok := m[c]
	if !ok {
		return color.RGBA{}
	}
	return r
}

// modelRectangle holds the corners of a canvas in model-space.
type modelRectangle struct {
	tl, br [2]float64
}

// canvas holds an image reference, a description of the model space, and constants to scale between pixel- and model-space.
type canvas struct {
	img            *image.RGBA
	modelRect      modelRectangle
	xscale, yscale float64
}

// NewCanvas generates a canvas and calculates constants to scale between pixel- and model-space.
func NewCanvas(w, h int, xmin, xmax, ymin, ymax float64) canvas {
	m := image.NewRGBA(image.Rect(0, 0, w, h))
	draw.Draw(m, m.Bounds(), &image.Uniform{DefaultColor("white")}, image.ZP, draw.Src)

	mr := modelRectangle{[2]float64{xmin, ymin}, [2]float64{xmax, ymax}}

	xrange := xmax - xmin
	yrange := ymax - ymin

	xscale := float64(w) / xrange
	yscale := float64(h) / yrange

	return canvas{m, mr, xscale, yscale}
}

// ExportToPNG writes a canvas to file.
func (C *canvas) ExportToPNG(filename string) {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fn := filepath.Join(wd, filename)

	f, err := os.Create(fn)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = png.Encode(f, C.img)
	if err != nil {
		panic(err)
	}
}

// Circle logic from the Go blog:
// https://blog.golang.org/go-imagedraw-package

type circle struct {
	o image.Point
	r int
}

func (c *circle) ColorModel() color.Model {
	return color.AlphaModel
}

func (c *circle) Bounds() image.Rectangle {
	return image.Rect(c.o.X-c.r, c.o.Y-c.r, c.o.X+c.r, c.o.Y+c.r)
}

func (c *circle) At(x, y int) color.Color {
	xx, yy, rr := float64(x-c.o.X)+0.5, float64(y-c.o.Y)+0.5, float64(c.r)
	if xx*xx+yy*yy < rr*rr {
		return color.Alpha{255}
	}
	return color.Alpha{0}
}

// DrawCircleAt draws a circle on a canvas. x/y are in model-space, r is in pixel-space.
func (C *canvas) DrawCircleAt(x, y float64, r int, fillColor color.RGBA) {
	X, Y := int((x-C.modelRect.tl[0])*C.xscale), int((y-C.modelRect.tl[1])*C.yscale)
	c := circle{image.Point{X, C.img.Bounds().Max.Y - Y}, r}
	draw.DrawMask(C.img, C.img.Bounds(), &image.Uniform{fillColor}, image.ZP, &c, image.ZP, draw.Over)
}
