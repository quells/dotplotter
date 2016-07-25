# dotplotter
Go package for drawing dots.

## Usage

`c := NewCanvas(w, h, xmin, xmax, ymin, ymax)`

Creates a canvas with an image of size (w, h) in pixels and a model space spanning from (xmin, ymin) to (xmax, ymax). The model space does not have to be 1:1 with the pixel space in the image, nor do the x-scale and y-scale need to be the same. Dots will always be drawn as circles, however; unequal axis scales will not result in ellipses.

`fillColor := DefaultColor("clr")`

A `map[string]color.RGBA` that provides RGB instances for common colors. Includes black, white, red, orange, yellow, green, blue, and purple. If the string is not in the map, then it returns a clear color.

`c.DrawCircleAt(x, y, r, fillColor)`

Draws a dot of radius r on the canvas' image centered at the position (x, y) in *model-space* filled with fillColor. Positive x-y (quadrant I) is in the upper-right.

`c.ExportToPNG("test.png")`

Saves the image to the current working directory.

## Example

![example result](https://github.com/quells/dotplotter/blob/master/test.png)

Generated with [example.go](https://github.com/quells/dotplotter/blob/master/example/example.go)
