package main

import (
	"testing"
)

type Rectangle struct {
	width, height int
	area          int
}

func (r *Rectangle) CalculateArea() {
	r.area = r.width * r.height
}

func (r *Rectangle) Area() int {
	return r.area
}

type Square struct {
	side int
	area int
}

func (s *Square) CalculateArea() {
	s.area = s.side * s.side
}

func (s *Square) Area() int {
	return s.area
}

// TestMain_fa0fc7ab4c is the test function for the main function
func TestMain_fa0fc7ab4c(t *testing.T) {
	// Test case for Rectangle
	rect := &Rectangle{width: 16, height: 6}
	rect.CalculateArea()
	if rect.Area() != 96 {
		t.Errorf("Expected area of rectangle to be 96, but got %d\n", rect.Area())
	}

	// Test case for Square
	sq := &Square{side: 8}
	sq.CalculateArea()
	if sq.Area() != 64 {
		t.Errorf("Expected area of square to be 64, but got %d\n", sq.Area())
	}
}
