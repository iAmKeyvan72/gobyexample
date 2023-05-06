package main

import (
	"fmt"
	"math"
)

type geometrics interface {
	area() float64
	perim() float64
}

type rect struct {
	width, height float64
}

type circle struct {
	radius float64
}

func (r *rect) area() float64 {
	return r.width * r.height
}
func (r *rect) perim() float64 {
	return 2 * (r.height + r.width)
}

func (c *circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c *circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometrics) {
	fmt.Println("area: ", g.area())
	fmt.Println("perim: ", g.perim())
}

func main() {
	rectangel := rect{height: 2, width: 5}
	measure(&rectangel)

	cir := circle{radius: 34}
	measure(&cir)
}
