package page7

import "math"

type shape interface {
	area() float64
	perimeter() float64
}
type rect struct {
	width, height float64
}
type circle struct {
	radius float64
}

func newRect(width, height float64) *rect {
	return &rect{width: width, height: height}
}

func (r rect) area() float64 {
	return r.width * r.height
}

func (r rect) perimeter() float64 {
	return 2 * (r.width + r.height)
}

func newCircle(radius float64) *circle {
	return &circle{radius: radius}
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}
