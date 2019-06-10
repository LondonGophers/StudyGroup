package surface

import (
	"fmt"
	"image/color"
	"math"
	"io"

	"github.com/billglover/gradient"
)

// Surface is a function that takes an x, y coordinate and returns the 
// corresponding z coordinate.
type Surface func(float64, float64) float64

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges [-xyrange..+xyrange]
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos (30°)

// Render takes an io.Writer and renders and SVG
func Render(w io.Writer, shape Surface) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)

	if shape == nil {
		shape = Original
	}	

	minZ, maxZ := getZRange(shape)

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, oka := corner(i+1, j, shape)
			bx, by, okb := corner(i, j, shape)
			cx, cy, okc := corner(i, j+1, shape)
			dx, dy, okd := corner(i+1, j+1, shape)
			if !(oka && okb && okc && okd) {
				continue
			}

			c := cellColor(i, j, minZ, maxZ, shape)

			fmt.Fprintf(w, "\t<polygon stroke='%s' fill='%s' fill-opacity='0.7' stroke-width='.25' points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B),
				fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B),
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
}

func cellColor(i, j int, minZ, maxZ float64, f Surface) color.RGBA {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	z := f(x, y)

	p := z - minZ/(maxZ-minZ)

	red := color.RGBA{0xFF, 0x00, 0x00, 0xFF}
	blue := color.RGBA{0x00, 0x00, 0xFF, 0xFF}
	c := gradient.LinearPoint(blue, red, p)
	return c
}

func getZRange(f func(float64, float64) float64) (float64, float64) {
	minZ := math.MaxFloat64
	maxZ := -math.MaxFloat64

	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {

			x := xyrange * (float64(i)/cells - 0.5)
			y := xyrange * (float64(j)/cells - 0.5)
			z := f(x, y)
			if z > maxZ {
				maxZ = z
			}
			if z < minZ {
				minZ = z
			}
		}
	}

	return minZ, maxZ
}

func corner(i, j int, f func(float64, float64) float64) (float64, float64, bool) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)
	ok := !(math.IsNaN(z) || math.IsInf(z, 0))

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, ok
}

// Original is the function for the original water-drop.
func Original(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}

// EggBox is the function for the egg-box.
func EggBox(x, y float64) float64 {
	r := math.Sin(x) - math.Sin(y)
	r = r / 10
	return r
}

// Saddle is the function for a saddle.
func Saddle(x, y float64) float64 {
	r := x*x - y*y
	r = r / 500
	return r
}

// MonkeySaddle is the function for a monkey-saddle.
func MonkeySaddle(x, y float64) float64 {
	r := 0.5*x*x*x - 3*x*y*y
	r = r / 10000
	return r
}

// Paper is the function for a sheet of paper.
func Paper(x, y float64) float64 {
	r := x*x*x - y*y*y
	r = r / 10000
	return r
}
