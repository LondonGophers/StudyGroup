// Surface computes an SVG rendering of a 3-D surface function.
package main

import (
	"errors"
	"fmt"
	"io"
	"math"
	"net/http"
)

func poly(out io.Writer, r *http.Request) {
	params := parseParams(r)

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", int(params["width"]), int(params["height"]))

	for i := 0; i < int(params["cells"]); i++ {
		for j := 0; j < int(params["cells"]); j++ {
			ax, ay, z, errA := corner(i+1, j, params)
			bx, by, _, errB := corner(i, j, params)
			cx, cy, _, errC := corner(i, j+1, params)
			dx, dy, _, errD := corner(i+1, j+1, params)

			if errA != nil || errB != nil || errC != nil || errD != nil {
				continue
			}

			fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: ",
				ax, ay, bx, by, cx, cy, dx, dy)

			switch {
			case z > .1:
				fmt.Fprint(out, "red")
			case z < -.1:
				fmt.Fprint(out, "blue")
			}

			fmt.Fprint(out, "'/>\n")
		}
	}
	fmt.Fprintln(out, "</svg>")
}

func corner(i, j int, params map[string]float64) (float64, float64, float64, error) {
	// Find point (x,y) at corner of cell (i,j).
	x := params["xyrange"] * (float64(i)/params["cells"] - 0.5)
	y := params["xyrange"] * (float64(j)/params["cells"] - 0.5)

	// Compute surface height z.
	z := f(x, y)

	if math.IsInf(z, 0) || math.IsNaN(z) {
		return 0, 0, 0, errors.New("non-finite value calculated")
	}

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := params["width"]/2 + (x-y)*params["cos30"]*params["xyscale"]
	sy := params["height"]/2 + (x+y)*params["sin30"]*params["xyscale"] - z*params["zscale"]

	return sx, sy, z, nil
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
