package main

import (
	"fmt"
	"math"
)

const (
	width, height = 600, 300
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6 // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	fmt.Printf("<svg xmlns='http://wwww.w3.org/2000/svg' "+
		"style='stroke: grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, ok1 := corner(i+1, j)
			bx, by, ok2 := corner(i, j)
			cx, cy, ok3 := corner(i, j+1)
			dx, dy, ok4 := corner(i+1, j+1)
			if (ok1 && ok2 && ok3 && ok4) {
			fmt.Printf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
				}
		}
	}
	fmt.Println("</svg>")
}

func corner(i, j int) (float64, float64, bool) {
	x := xyrange * (float64(i)/(cells-1) - 0.5)
	y := xyrange * (float64(j)/(cells-1) - 0.5)

	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, false
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, true
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}