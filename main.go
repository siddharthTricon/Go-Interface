package main

import (
	"fmt"
)

type Rectangle struct{
	Height float64
	Width float64
}

type Circle struct{
	Radius float64
}

type Shape interface{
	Area() float64
	Perimeter() float64
}

func (r Rectangle) Area() float64{
	return (r.Height * r.Width)
}

func (r Rectangle) Perimeter() float64{
	return 2*(r.Height * r.Width)
}

func (c Circle) Area() float64{
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64{
	return 4 * 3.14 * c.Radius
}

func PrintShape(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main (){
	rect := Rectangle{Height: 5, Width: 6}
	PrintShape(rect)

	circle := Circle{Radius: 4}
	PrintShape((circle))
}