package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
)

// Colours from http://www.flatuicolorpicker.com

var Green = color.RGBA{0, 230, 64, 1}
var Orange = color.RGBA{230, 126, 34, 1}
var Purple = color.RGBA{154, 18, 179, 1}
var Yellow = color.RGBA{240, 255, 0, 1}
var Red = color.RGBA{219, 10, 91, 1}
var Blue = color.RGBA{77, 19, 209, 1}
var Black = color.Black
var White = color.White

var palette = []color.Color{Black, Green, Orange, Purple, Yellow, Red, Blue, White}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Int()%7+1))
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}