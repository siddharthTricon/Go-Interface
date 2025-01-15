package main

import (
	"fmt"
)

type Rectangle struct{
	Height float64
	Width float64
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

func PrintShape(s Shape) {
	fmt.Printf("Area: %f\n", s.Area())
	fmt.Printf("Perimeter: %f\n", s.Perimeter())
}

func main (){
	rect := Rectangle{Height: 5, Width: 6}
	PrintShape(rect)
}