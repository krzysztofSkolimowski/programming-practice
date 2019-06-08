package intersection

import (
	"github.com/pkg/errors"
	"math"
)

var (
	ErrPointsAreEqual = errors.New("Given points are equal")
)

type Point struct {
	X, Y float64
}

func (p Point) equals(another Point) bool { return (p.X == another.X) && (p.Y == another.Y) }

type line struct {
	start, end        Point
	slope, interceptY float64
}

func NewLine(start, end Point) (line, error) {
	if start.equals(end) {
		return line{}, ErrPointsAreEqual
	}

	if start.X > end.X {
		start, end = end, start
	}

	slope := math.Atan((start.Y - end.Y) / (start.X - end.X))
	interceptY := end.Y - slope*end.X

	return line{start, end, slope, interceptY}, nil
}

func (l line) Parallel(l2 line) bool {
	return l.slope == l2.slope
}

func Intersect(a, b line) bool {
	if a.Parallel(b) {
		return false
	}

	x := (b.interceptY - a.interceptY) / (a.slope - b.slope)
	intersection := Point{x, a.slope*x + a.interceptY}
	return intersection.X > a.start.X && intersection.X < a.end.X
}
