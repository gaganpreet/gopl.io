package geometry

import "math"

type Point struct{ X, Y float64 }

type Path []Point

// traditional
func Distance(p, q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y-p.Y)
}

// method of Point type
func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X - p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(x float64) {
	p.X *= x
	p.Y *= x
}

func (p *Point) ScaleByPointer(x float64) {
	(*p).X *= x
	(*p).Y *= x
}

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}

type IntList struct {
	Value int
	Tail *IntList
}

func (l *IntList) Sum() int {
	if l == nil {
		return 0
	}
	return l.Value + l.Tail.Sum()
}
