package main

import (
	"testing"
)

type Rectangle struct {
	width, height float64
}

func (r *Rectangle) CalculateArea() float64 {
	return r.width * r.height
}

type Square struct {
	side float64
}

func (s *Square) CalculateArea() float64 {
	return s.side * s.side
}

func TestRectangleArea(t *testing.T) {
	rect := &Rectangle{width: 16, height: 6}
	area := rect.CalculateArea()
	if area != 96 {
		t.Errorf("Area of rectangle was incorrect, got: %f, want: %f.", area, 96.0)
	}
}

func TestSquareArea(t *testing.T) {
	sq := &Square{side: 8}
	area := sq.CalculateArea()
	if area != 64 {
		t.Errorf("Area of square was incorrect, got: %f, want: %f.", area, 64.0)
	}
}
