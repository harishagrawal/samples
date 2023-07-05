package main

import (
	"testing"
)

type Rectangle struct {
	width, height float64
}

func (rect *Rectangle) CalculateArea() float64 {
	return rect.width * rect.height //area of rectangle
}

func TestCalculateArea_acb586a213(t *testing.T) {
	// Test case 1: Normal scenario
	rect1 := &Rectangle{width: 10, height: 5}
	area1 := rect1.CalculateArea()
	if area1 != 50 {
		t.Errorf("Failed! Expected area: %v, got: %v", 50, area1)
	}

	// Test case 2: Zero width and height
	rect2 := &Rectangle{width: 0, height: 0}
	area2 := rect2.CalculateArea()
	if area2 != 0 {
		t.Errorf("Failed! Expected area: %v, got: %v", 0, area2)
	}

	// Test case 3: Negative width and height
	rect3 := &Rectangle{width: -10, height: -5}
	area3 := rect3.CalculateArea()
	if area3 != 50 {
		t.Errorf("Failed! Expected area: %v, got: %v", 50, area3)
	}

	// Test case 4: One side is negative
	rect4 := &Rectangle{width: -10, height: 5}
	area4 := rect4.CalculateArea()
	if area4 != -50 {
		t.Errorf("Failed! Expected area: %v, got: %v", -50, area4)
	}
}
