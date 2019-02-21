package main

import (
	"flag"
	"fmt"
	"image"
	"image/color"
	"image/color/palette"
	"image/gif"
	"io"
	"log"
	"math"
	"math/big"
	"math/cmplx"
	"net/http"
	"os"
	"strconv"
)

var v = flag.Bool("v", false, "print verbose logs for debugging")
var c = flag.String("c", "websafe", "show in colours")
var exp = flag.Float64("s", 1.414, "zoom")
var left = flag.Float64("l", 0, "left shift")
var down = flag.Float64("d", 0, "down shift")
var fractal = flag.String("f", "mandel", "fractal type")
var nFrames = flag.Int("i", 128, "number of frames")
var web = flag.Bool("web", false, "web server")
var crl = flag.Float64("real", 0.4245127190500396, "center real")
var cim = flag.Float64("imag", 0.2075302281667453, "center imag")

type Grey = color.Gray

var roots = make(map[complex128]int)

var key = -1
var colours = [4]color.RGBA{
	color.RGBA{255, 0, 0, 255},
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 0, 255, 255},
	color.RGBA{255, 255, 0, 255},
}

var (
	cntr = 0.4245127190500396 + 0.2075302281667453i // precise enough
)

const (
	//cntr = 0 + 0i
	//		cntr          = -1.45 - 0i
	//https://commons.wikimedia.org/wiki/File:Zoom_around_principal_Misiurewicz_point_for_periods_from_2_to_1024.gif
	// https://en.wikipedia.org/wiki/Misiurewicz_point
	//cntr = -0.77568377 + 0.13646737i // very pretty - but not precise enough
	//cntr = 0.3663629834227643 + 0.5915337732614452i // precise enough
	//cntr = 0.4245127190500396 + 0.2075302281667453i // precise enough
	//cntr          = -1.54368901269109 + 0i
	//cntr          = -0.260143 + 0.6337i //pretty not precise enough
	//cntr = 0.2759353624416824 + 0.0069166138017372i //very pretty buf goes down a hole
	//cntr          = 0.2787724591293833 + 0.0081245796484104i //very pretty buf goes down a hole
	//cntr = 0.3115076602815077 + 0.0271737195013418i //very pretty buf goes down a hole
	//	cntr          = -0.1010963638456221 + 0.9562865108091415i // precise enough
	//cntr          = 0.2501502296489224 + 0.0000029308049747i // nice but goes off side
	//cntr          = 0.4379242413594628 + 0.3418920843381161i //precise enough
	//cntr = -1i // good and precise enough so lets use this one
	//cntr          = -0.75 + 0i
	//cntr          = -1.401155 + 0i
	//cntr = -0.75 + 0.1i // sea horse valley
	//cntr          = 0.3 + 0.1i // elephant valley
	ymin, ymax    = -2, +2
	xmin, xmax    = -2, +2
	width, height = 256, 256

	delay = 16 //delay between frames in 10ms units
)

func main() {

	flag.Parse()
	cntr = complex(*crl, *cim)
	if *web {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			rl := r.FormValue("rl")
			mg := r.FormValue("mg")
			zoom := r.FormValue("zoom")
			if v, err := strconv.ParseFloat(rl, 64); err == nil {
				*crl = v
			}
			if v, err := strconv.ParseFloat(mg, 64); err == nil {
				*cim = v
			}
			if v, err := strconv.ParseFloat(zoom, 64); err == nil {
				*exp = v
			}
			cntr = complex(*crl, *cim)
			render(w)
		})
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}
	render(os.Stdout)
	//png.Encode(os.Stdout, img)
}

func render(w io.Writer) {

	anim := gif.GIF{LoopCount: *nFrames}
	rect := image.Rect(0, 0, width, height)

	//
	scale := 1.0
	for i := 0; i < *nFrames; i++ {
		//log.Println(scale)
		img := image.NewPaletted(rect, palette.WebSafe)
		for py := 0; py < height; py++ {
			// project py onto domain using y
			y := (float64(py)/height*(ymax-ymin))/(scale) + ymin/(scale)
			for px := 0; px < width; px++ {
				x := float64(px)/width*(xmax-xmin)/(scale) + xmin/(scale)
				z := complex(x, y)
				//
				switch *fractal {
				case "mandel":
					img.SetColorIndex(px, py, mandelbrot128Index(cntr+z))
				default:

					img.SetColorIndex(px, py, newtonIndex(cntr+z))
				}
			}
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
		scale = scale * (*exp)
		//log.Println(scale, *exp)
	}
	//

	verbosef("# of roots: %d\n", len(roots))
	if *v {
		for k := range roots {
			verbosef("root is: %v\n", k)
		}
	}
	gif.EncodeAll(w, &anim)

}

// next guess is z - ( f(x) / 'f(x) )
func newton(z complex128) color.Color {

	const (
		iter     = 50
		contrast = 20
		dz       = 0.000001
		rnd      = 6
	)

	// next guess is z - ( f(x) / 'f(x) )
	for i := 0; i < iter; i++ {
		nextGuess := (cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3))
		z = z - nextGuess

		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 0+dz {
			if _, ok := roots[round(z, rnd)]; !ok {
				key++
				roots[round(z, rnd)] = key
			}

			switch *c {
			case "grey":
				return Grey{255 - uint8(i)*contrast}
			case "websafe":
				return palette.WebSafe[(255-uint(i)*contrast)%216]
			case "plan9":
				return palette.Plan9[(255-uint(i)*contrast)%255]
			case "primary":
				return shaded(colours[roots[round(z, rnd)]], i, contrast)
			default:
				return palette.WebSafe[(255-uint(i)*contrast)%216]
			}
		}
	}
	return color.Black

}

// next guess is z - ( f(x) / 'f(x) )
func newtonIndex(z complex128) uint8 {

	const (
		iter     = 50
		contrast = 20
		dz       = 0.000001
		rnd      = 6
	)

	// next guess is z - ( f(x) / 'f(x) )
	for i := 0; i < iter; i++ {
		nextGuess := (cmplx.Pow(z, 4) - 1) / (4 * cmplx.Pow(z, 3))
		z = z - nextGuess

		if cmplx.Abs(cmplx.Pow(z, 4)-1) < 0+dz {
			if _, ok := roots[round(z, rnd)]; !ok {
				key++
				roots[round(z, rnd)] = key
			}

			//			switch *c {
			//			case "grey":
			//				return Grey{255 - uint8(i)*contrast}
			//			case "websafe":
			//				return palette.WebSafe[(255-uint(i)*contrast)%216]
			//			case "plan9":
			//				return palette.Plan9[(255-uint(i)*contrast)%255]
			//			case "primary":
			//				return shaded(colours[roots[round(z, rnd)]], i, contrast)
			//			default:
			//				return palette.WebSafe[(255-uint(i)*contrast)%216]
			//			}
			return uint8((255 - uint(i)*contrast) % 216)

		}
	}
	return 0

}

func shaded(c color.Color, i int, contrast int) color.Color {
	r, g, b, _ := c.RGBA()
	if r != 0 {
		r = uint32(255-(i*contrast)) % 255
	}
	if g != 0 {
		g = uint32(255-(i*contrast)) % 255
	}
	if b != 0 {
		b = uint32(255-(i*contrast)) % 255
	}
	return color.RGBA{uint8(r), uint8(g), uint8(b), 255}
}

func round(z complex128, r int) complex128 {
	rl := real(z)
	mg := imag(z)

	rl = math.Round(rl * math.Pow10(r))
	mg = math.Round(mg * math.Pow10(r))

	rl = rl / math.Pow10(r)
	mg = mg / math.Pow10(r)

	return complex(rl, mg)
}

func verbosef(format string, args ...interface{}) {
	if *v {
		fmt.Fprintf(os.Stderr, format, args...)
	}
}

func mandelbrot128(z complex128) color.Color {
	const iterations = 500
	const contrast = 80
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if c := cmplx.Abs(v); c > 2 {
			return palette.WebSafe[(255-uint(n)*contrast)%216]
		}
	}
	return color.Black
}

func mandelbrot128Index(z complex128) uint8 {
	const iterations = 500
	const contrast = 80
	var v complex128
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if c := cmplx.Abs(v); c > 2 {
			return uint8((255 - uint(n)*contrast) % 216)
		}
	}
	return 0
}

func mandelbrot64(z complex64) color.Color {
	const iterations = 500
	const contrast = 80
	var v complex64
	for n := 0; n < iterations; n++ {
		v = v*v + z
		if c := cmplx.Abs(complex128(v)); c > 2 {
			return palette.WebSafe[(255-uint(n)*contrast)%216]

		}
	}
	return color.Black
}

func mandelbrotFloat64(rl float64, mg float64) color.Color {
	const iterations = 500
	const contrast = 80
	var vrl float64
	var vmg float64

	for n := 0; n < iterations; n++ {
		vrl, vmg = (vrl*vrl-vmg*vmg)+rl, (vrl*vmg+vrl*vmg)+mg
		if c := math.Sqrt(vrl*vrl + vmg*vmg); c > 2 {
			return palette.WebSafe[(255-uint(n)*contrast)%216]
		}
	}
	return color.Black
}

func mandelbrotBigFloat(rl *big.Float, mg *big.Float) color.Color {

	const iterations = 500
	const contrast = 80
	var vrl, vmg = big.NewFloat(0), big.NewFloat(0)

	for n := 0; n < iterations; n++ {
		//		vrl, vmg = (vrl*vrl-vmg*vmg)+rl, (vrl*vmg+vrl*vmg)+mg
		a, b, c, d := big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), big.NewFloat(0)
		//log.Println(a, b, c, d)
		a = a.Mul(vrl, vrl)
		b = b.Mul(vmg, vmg)
		c = c.Sub(a, b)
		d = d.Add(c, rl)

		e, f, g := big.NewFloat(0), big.NewFloat(0), big.NewFloat(0)
		e = e.Mul(vrl, vmg)
		f = f.Add(e, e)
		g = g.Add(f, mg)

		vrl = d
		vmg = g

		i, j, k, l := big.NewFloat(0), big.NewFloat(0), big.NewFloat(0), big.NewFloat(4)
		i = i.Mul(vrl, vrl)
		j = j.Mul(vmg, vmg)
		k = k.Add(i, j)
		//log.Println(i, j, k, l, k.Cmp(l))
		if k.Cmp(l) > 0 {
			return palette.WebSafe[(255-uint(n)*contrast)%216]
		}
	}
	return color.Black
}

var cnt = 0

//too slow... dont use!
func mandelbrotBigRat(rl *big.Rat, mg *big.Rat) color.Color {
	const iterations = 500
	const contrast = 80
	var vrl, vmg = big.NewRat(0, 1), big.NewRat(0, 1)

	cnt++

	log.Println("cnt", cnt)
	for n := 0; n < iterations; n++ {
		if cnt == 4028 {
			log.Println("n", n)
		}
		a, b, c, d := big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(0, 1)
		a = a.Mul(vrl, vrl)
		b = b.Mul(vmg, vmg)
		c = c.Sub(a, b)
		d = d.Add(c, rl)

		e, f, g := big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(0, 1)
		e = e.Mul(vrl, vmg)
		f = f.Add(e, e)
		g = g.Add(f, mg)

		vrl = d
		vmg = g

		i, j, k, l := big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(0, 1), big.NewRat(4, 1)
		i = i.Mul(vrl, vrl)
		j = j.Mul(vmg, vmg)
		k = k.Add(i, j)
		if k.Cmp(l) > 0 {
			return palette.WebSafe[(255-uint(n)*contrast)%216]
		}
	}
	return color.Black
}
