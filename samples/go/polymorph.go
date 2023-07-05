package main

import "fmt"

type Shape struct {
	area float64 //create area field in the struct
}

func (sq *Shape) Area() float64 {
	fmt.Println("Calling Shape Area")
	return sq.area //return area of square
}

type Rectangle struct {
	Shape
	width  float64 //width of rectangle
	height float64 //height of rectangle
}

func (rect *Rectangle) CalculateArea() {
	rect.area = rect.width * rect.height //area of rectangle
}

type Square struct {
	Shape
	side float64 //side of square
}

func (sq *Square) CalculateArea() {
	sq.area = sq.side * sq.side //area of square
}

func main() {
	rect := &Rectangle{width: 16, height: 6} //set the width and height of square
	rect.CalculateArea()
	fmt.Println("Area of rectangle: ", rect.Area()) //print area of rectangle

	sq := &Square{side: 8} //set side of square
	sq.CalculateArea()
	fmt.Println("Area of square: ", sq.Area()) //print area of square.
}
