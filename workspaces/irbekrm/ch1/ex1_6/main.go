package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

const (
	backgrounColorIndex = iota // first color in palette
	patternColorIndex          // next color in palette
)

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5     // number of complete x oscillator revolutions
		res     = 0.001 // angular resolution
		size    = 100
		nframes = 64
		delay   = 8 // delay betwen frames in 10ms units
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		palette := makePalette()
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), patternColorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}

func makePalette() []color.Color {
	c := randomColor()
	palette := []color.Color{color.Black, c}
	return palette
}
func randomColor() color.RGBA {
	tempColors := []uint8{}
	for i := 0; i < 4; i++ {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		randInt := uint8(r.Intn(255))
		tempColors = append(tempColors, randInt)
	}
	return color.RGBA{tempColors[0], tempColors[1], tempColors[2], tempColors[3]}
}
