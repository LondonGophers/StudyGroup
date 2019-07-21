// Following the approach of the Lissajous example in Section 1.7, construct a web
// server that computes surfaces and writes SVG data to the client. The server must
// set the `Content-Type` header like this:
// ```
// w.Header().Set("Content-Type", "image/svg+xml")
// ```
// (This step was not required in the Lissajous example because the server uses
// standard heuristics to recognize common formats like PNG from the first 512
// bytes of the response, and generates the proper header.) Allow the client to
// specify values like height, width, and color as HTTP request parameters.
package main

import (
	"errors"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // axis ranges (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
	angle         = math.Pi / 6         // angle of x, y axes (=30°)
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle) // sin(30°), cos(30°)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		plot(w)
	}

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}

func plot(out io.Writer) {
	result := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, err1 := corner(i+1, j)
			bx, by, err2 := corner(i, j)
			cx, cy, err3 := corner(i, j+1)
			dx, dy, err4 := corner(i+1, j+1)

			if err1 == nil && err2 == nil && err3 == nil && err4 == nil {
				result += fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
					ax, ay, bx, by, cx, cy, dx, dy)
			}
		}
	}
	result += fmt.Sprintf("</svg>")
	io.WriteString(out, result)
}

func corner(i, j int) (float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsNaN(z) || math.IsInf(z, 0) || math.IsInf(z, -1) {
		return 0, 0, errors.New("Non-finite number calculated")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
