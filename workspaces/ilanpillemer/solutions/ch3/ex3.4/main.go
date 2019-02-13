package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"os"
	"strconv"
)

var (
	height float64 = 300
	zscale float64 = float64(height) * 0.4
)

const (
	width   = 600
	cells   = 100
	xyrange = 30.0
	zrange  = 70.0
	xyscale = width / 2 / xyrange

	angle = math.Pi / 6 // 30 degrees
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)
var eq = flag.String("eq", "orig", "orig or monkey or egg")
var web = flag.Bool("server", false, "run as web server")

func main() {
	flag.Parse()
	f := toFunc(*eq)
	z := zrange
	if *web {
		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			r.ParseForm()
			if r.FormValue("eq") != "" {
				f = toFunc(r.FormValue("eq"))
			}
			if r.FormValue("zscale") != "" {
				z, _ = strconv.ParseFloat(r.FormValue("zscale"), 64) //ignoring errors
			}
			if r.FormValue("height") != "" {
				height, _ = strconv.ParseFloat(r.FormValue("height"), 64) //ignoring errors
			}

			w.Header().Set("Content-Type", "image/svg+xml")
			render(w, z, f)
		})
		log.Fatal(http.ListenAndServe("localhost:8080", nil))
	}
	out := os.Stdout
	render(out, z, f)
}

func toFunc(n string) (f func(x, y float64) float64) {
	switch n {
	case "well":
		return well
	case "circle":
		return circle
	case "egg":
		return egg
	case "monkey":
		return monkeySaddle
	case "sin":
		return sin
		case "saddle":
		return saddle
	default:
		return orig
	}
}

func render(out io.Writer, z float64, f func(x, y float64) float64) {

	fmt.Fprintf(out, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill:white; stroke-width:0.7' "+
		"width='%d' height='%d'>", width, int(height))
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, colour, ok1 := corner(i+1, j, z, f)
			bx, by, colour, ok2 := corner(i, j, z, f)
			cx, cy, colour, ok3 := corner(i, j+1, z, f)
			dx, dy, colour, ok4 := corner(i+1, j+1, z, f)
			if ok1 && ok2 && ok3 && ok4 {
				fmt.Fprintf(out, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='fill:%s' />\n",
					ax, ay, bx, by, cx, cy, dx, dy, colour)
			}
		}
	}
	fmt.Fprintln(out, "</svg>")
}

//zrange should shadow constant.. icky code..
func corner(i int, j int, zrange float64, f func(x, y float64) float64) (float64, float64, string, bool) {
	colour := "#00ff00"
	x := xyrange * (float64(i)/(cells-1) - 0.5)
	y := xyrange * (float64(j)/(cells-1) - 0.5)

	z := math.Mod(f(x, y), float64(height))
	z = zrange * (float64(z) / (cells - 1))

	if math.IsNaN(z) || math.IsInf(z, 0) {
		return 0, 0, colour, false
	}

	if z > 0.1 {
		colour = "#ff0000"
	}

	if z < 0.005 {
		colour = "#0000ff"
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := float64(height)/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, colour, true
}

func orig(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

func well(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r * math.Sin(x)
}

func saddle(x, y float64) float64 {
	z := math.Pow(x, 2) - math.Pow(y, 2)
	return z
}

func egg(x, y float64) float64 {
	z := math.Sin(x) * math.Sin(y)
	return z
}

func sin(x, y float64) float64 {
	z := math.Sin(x)
	return z
}

func circle(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Pi * r * r
}

func monkeySaddle(x, y float64) float64 {
	z := math.Pow(x, 3) - 3*x*math.Pow(y, 2)
	return z
}

