package main

import (
	"fmt"
	"math"
)

type Circle struct {
	Radius float64
}

type Rectange struct {
	Height float64
	Width  float64
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func (c *Circle) Area() float64 {
	return c.Radius * c.Radius * math.Pi
}

func (c *Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func (r *Rectange) Area() float64 {
	return r.Height * r.Width
}

func (r *Rectange) Perimeter() float64 {
	return (r.Height + r.Width) * 2
}

func PrintArea(s Shape) {
	fmt.Println("Area -> ", s.Area())
}

func PrintPerimeter(s Shape) {
	fmt.Println("Perimeter -> ", s.Perimeter())
}

func main() {
	circle := &Circle{Radius: 3}
	rectangle := &Rectange{Height: 2, Width: 4}
	PrintArea(circle)
	PrintPerimeter(circle)
	PrintArea(rectangle)
	PrintPerimeter(rectangle)
}
