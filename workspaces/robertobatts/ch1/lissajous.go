package main

import (
	"bufio"
	"fmt"
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"os"
)

func createGif() {
	file, _ := os.Create("lissajous.gif")
	out := bufio.NewWriter(file)
	lissajous(out)
}

func createPalette() []color.Color {
	var palette = []color.Color{color.Black}
	a := uint8(1)
	for i := uint8(0); i < 254; i++ {
		palette = append(palette, color.RGBA{a*3 - i, a + i*2, i + 17, i})
	}
	return palette
}

func lissajous(out io.Writer) {
	palette := createPalette()
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.002 // angular resolution
		size    = 200   // image canvas covers [size..+ size]
		nframes = 64    // number of animation frames
		delay   = 8     // delay between frames in 10ms units
	)
	freq := 1.1
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		fmt.Printf("Frame: %d\n", i)
		rect := image.Rect(0, 0, 2*size, 2*size)
		img := image.NewPaletted(rect, palette)
		colorIndex := uint8(1)
		for t := 0.0; t < 2*math.Pi*cycles; t += res {
			x := math.Sin(t)
			y := math.Sin(freq*t + phase)
			img.SetColorIndex(size+int(size*x), size+int(size*y), colorIndex)
			colorIndex++
			if colorIndex == 0 {
				colorIndex++
			}
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
