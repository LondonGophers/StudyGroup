/*
Package image implements functionality to build an SVG image
*/
package image

import (
	"fmt"
	"math"
	"strings"
)

type Client struct {
	width float64
	height float64
	cells int
	xyrange float64
	xyscale float64
	zscale float64
	sin30 float64
	cos30 float64
	color string
}

func New(height, width float64, color string) *Client {
	c :=  Client{
		width: width,
		height: height,
		cells: 100,
		xyrange: 30.0,
		color: color,
	}
	c.xyscale = c.width / 2 / c.xyrange
	c.zscale = c.height * 0.4
	angle := math.Pi / 6
	c.sin30 = math.Sin(angle)
	c.cos30 = math.Cos(angle)
	return &c
}

func (c *Client) BuildImage() string {
	var sb strings.Builder
	header := fmt.Sprintf("<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: %s; stroke-width: 0.7' "+
		"width='%d' height='%d'>", c.color, c.width, c.height)
	sb.WriteString(header)
	for i := 0; i < c.cells; i++ {
		for j := 0; j < c.cells; j++ {
			ax, ay := c.corner(i+1, j)
			bx, by := c.corner(i, j)
			cx, cy := c.corner(i, j+1)
			dx, dy := c.corner(i+1, j+1)
			p := fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' />\n",
				ax, ay, bx, by, cx, cy, dx, dy)
			sb.WriteString(p)
		}
	}
	sb.WriteString("</svg>")
	return sb.String()
}

func (c *Client) corner(i, j int) (float64, float64) {
	// Find point (x, y) at corner of cell (i, j).
	x := c.xyrange * (float64(i)/float64(c.cells) - 0.5)
	y := c.xyrange * (float64(j)/float64(c.cells) - 0.5)

	// Compute surface height z
	z := f(x, y)

	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := c.width/2 + (x-y)*c.cos30*c.xyscale
	sy := c.height/2 + (x+y)*c.sin30*c.xyscale - z*c.zscale
	return sx, sy
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	return math.Sin(r) / r
}
