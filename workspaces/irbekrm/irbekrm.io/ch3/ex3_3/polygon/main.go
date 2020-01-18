package polygon

import "fmt"

import "math"

import "strings"

const (
	MIN_COLOR_STRENGTH float64 = 0.0
	MAX_COLOR_STRENGTH float64 = 255.0
)

type polygon struct {
	ax       float64
	ay       float64
	bx       float64
	by       float64
	cx       float64
	cy       float64
	dx       float64
	dy       float64
	height   float64
	redness  int
	blueness int
}
type PolygonsCollection struct {
	currentMinHeight float64
	currentMaxHeight float64
	polygons         []*polygon
}

func NewPolygonsCollection() *PolygonsCollection {
	return &PolygonsCollection{
		currentMinHeight: math.Inf(1),
		currentMaxHeight: math.Inf(-1),
	}
}

func (t *PolygonsCollection) NewPolygon(ax, ay, bx, by, cx, cy, dx, dy, height float64) {
	p := &polygon{
		ax:     ax,
		ay:     ay,
		bx:     bx,
		by:     by,
		cx:     cx,
		cy:     cy,
		dx:     dx,
		dy:     dy,
		height: height,
	}
	if height > t.currentMaxHeight {
		t.currentMaxHeight = height
	}
	if height < t.currentMinHeight {
		t.currentMinHeight = height
	}
	t.polygons = append(t.polygons, p)
}

func (t *PolygonsCollection) GetSVG() string {
	var sb strings.Builder
	for _, polygon := range t.polygons {
		polygon.setColors(t.currentMaxHeight, t.currentMinHeight)
		p := polygon.toSVG()
		sb.WriteString(p)
	}
	return sb.String()
}

func (p *polygon) setColors(maxHeight, minHeight float64) {
	p.setRedness(maxHeight, minHeight)
	p.setBlueness(maxHeight, minHeight)
}

func (p *polygon) setRedness(maxHeight, minHeight float64) {
	r := (p.height - minHeight) * (MAX_COLOR_STRENGTH - MIN_COLOR_STRENGTH) / (maxHeight - minHeight)
	// fmt.Printf("Min height is %v, max height is %v, this height is %v\nredness is %v\n", minHeight, maxHeight, p.height, r)
	p.redness = int(r)
}

// https://stackoverflow.com/questions/13137348/scaling-range-of-values-with-negative-numbers
func (p *polygon) setBlueness(maxHeight, minHeight float64) {
	b := (p.height - maxHeight) * (MAX_COLOR_STRENGTH - MIN_COLOR_STRENGTH) / (minHeight - maxHeight)
	fmt.Printf("Min height is %v, max height is %v, this height is %v\nblueness is %v\n", minHeight, maxHeight, p.height, b)
	p.blueness = int(b)
}

func (p *polygon) toSVG() string {
	return fmt.Sprintf("<polygon points='%g,%g %g,%g %g,%g %g,%g' fill='#%02X00%02X' />", p.ax, p.ay, p.bx, p.by, p.cx, p.cy, p.dx, p.dy, p.redness, p.blueness)
}
